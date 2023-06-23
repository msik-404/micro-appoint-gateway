package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpctohttp"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func GetCompany(c *gin.Context) {
	companyID := c.Param("company_id")
	if err := middleware.IsProperObjectIDHex(companyID); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var companiesConn *grpc.ClientConn
	companiesConn, err := grpc.Dial(companies.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer companiesConn.Close()
	companiesClient := companiespb.NewApiClient(companiesConn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	companiesMessage := companiespb.CompanyRequest{Id: &companyID}
	companiesReply, err := companiesClient.FindOneCompany(ctx, &companiesMessage)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
		return
	}

	var employeesConn *grpc.ClientConn
	employeesConn, err = grpc.Dial(employees.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer employeesConn.Close()
	employeesClient := employeespb.NewApiClient(employeesConn)
	employeesMessage := employeespb.EmployeesRequest{CompanyId: &companyID}
	employeesReply, err := employeesClient.FindManyEmployees(ctx, &employeesMessage)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
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
		Name            *string                      `json:"name,omitempty"`
		Type            *string                      `json:"type,omitempty"`
		Localisation    *string                      `json:"localisation,omitempty"`
		LongDescription *string                      `json:"long_description,omitempty"`
		Services        []*companiespb.Service       `json:"services,omitempty"`
		Employees       []*employeespb.EmployeeShort `json:"employees,omitempty"`
	}
	response := Response{
		Name:            companiesReply.Name,
		Type:            companiesReply.Type,
		Localisation:    companiesReply.Localisation,
		LongDescription: companiesReply.LongDescription,
		Services:        companiesReply.Services,
		Employees:       employees,
	}
	c.JSON(http.StatusOK, response)
}
