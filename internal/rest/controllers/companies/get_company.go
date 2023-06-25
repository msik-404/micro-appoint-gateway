package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func GetCompany(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyID := c.Param("company_id")
		if err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		companiesClient := companiespb.NewApiClient(conns.GetCompaniesConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		companiesMessage := companiespb.CompanyRequest{Id: &companyID}
		companiesReply, err := companiesClient.FindOneCompany(ctx, &companiesMessage)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}

		employeesClient := employeespb.NewApiClient(conns.GetEmployeesConn())
		employeesMessage := employeespb.EmployeesRequest{CompanyId: &companyID}
		employeesReply, err := employeesClient.FindManyEmployees(ctx, &employeesMessage)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			if status != http.StatusNotFound {
				c.AbortWithError(status, err)
				return
			}
		}

		var employees []*employeespb.EmployeeShort
		if employeesReply != nil {
			employees = employeesReply.Employees
		}

		type Response struct {
			Name             *string                      `json:"name,omitempty"`
			Type             *string                      `json:"type,omitempty"`
			Localisation     *string                      `json:"localisation,omitempty"`
			ShortDescription *string                      `json:"short_description,omitempty"`
			LongDescription  *string                      `json:"long_description,omitempty"`
			Services         []*companiespb.Service       `json:"services,omitempty"`
			Employees        []*employeespb.EmployeeShort `json:"employees,omitempty"`
		}
		response := Response{
			Name:             companiesReply.Name,
			Type:             companiesReply.Type,
			Localisation:     companiesReply.Localisation,
			ShortDescription: companiesReply.ShortDescription,
			LongDescription:  companiesReply.LongDescription,
			Services:         companiesReply.Services,
			Employees:        employees,
		}
		c.JSON(http.StatusOK, response)
	}
}
