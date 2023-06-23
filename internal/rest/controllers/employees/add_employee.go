package employees

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpctohttp"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
	"github.com/msik-404/micro-appoint-gateway/internal/strtime"
)

func AddEmployee(c *gin.Context) {
	owner, err := middleware.GetOwner(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	companyID := c.Param("company_id")
	if err := middleware.IsProperObjectIDHex(companyID); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if _, ok := owner.Companies[companyID]; !ok {
		c.AbortWithError(
			http.StatusUnauthorized,
			errors.New("This owner does not own this company"),
		)
		return
	}
    
    newEmployee, err := middleware.GetData[middleware.Employee](c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := employeespb.AddEmployeeRequest{
		CompanyId: &companyID,
		Name:      newEmployee.Name,
		Surname:   newEmployee.Surname,
	}
	if newEmployee.WorkTimes != nil {
		workTimes, err := strtime.ToWorkTimes(newEmployee.WorkTimes)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		message.WorkTimes = workTimes
	}
	for _, hex := range newEmployee.Competence {
		if err := middleware.IsProperObjectIDHex(hex); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		message.Competence = append(message.Competence, hex)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(employees.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	client := employeespb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.AddEmployee(ctx, &message)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
        return
	}
	c.JSON(http.StatusOK, reply)
}
