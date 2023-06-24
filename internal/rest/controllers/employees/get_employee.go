package employees

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
	"github.com/msik-404/micro-appoint-gateway/internal/strtime"
)

func GetEmployee(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("employee_id")
		if err := middleware.IsProperObjectIDHex(employeeID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		client := employeespb.NewApiClient(conns.GetEmployeesConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		message := employeespb.EmployeeRequest{
			Id: &employeeID,
		}
		reply, err := client.FindOneEmployee(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
		type Employee struct {
			Name       *string               `json:"name,omitempty"`
			Surname    *string               `json:"surname,omitempty"`
			WorkTimes  *strtime.WorkTimesStr `json:"work_times,omitempty"`
			Competence []string              `json:"competence,omitempty"`
		}
		response := Employee{
			Name:       reply.Name,
			Surname:    reply.Surname,
			Competence: reply.Competence,
		}
		if reply.WorkTimes != nil {
			workTimesStr, err := strtime.ToWorkTimesStr(reply.WorkTimes)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			response.WorkTimes = &workTimesStr
		}
		c.JSON(http.StatusOK, response)
	}
}
