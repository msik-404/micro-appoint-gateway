package employees

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
	"github.com/msik-404/micro-appoint-gateway/internal/strtime"
)

func GetEmployee(c *gin.Context) {
	employeeID := c.Param("employee_id")
	if err := middleware.IsProperObjectIDHex(employeeID); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(employees.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	client := employeespb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := employeespb.EmployeeRequest{
		Id: &employeeID,
	}
	reply, err := client.FindOneEmployee(ctx, &message)

	if err != nil {
		code := status.Code(err)
		if code == codes.InvalidArgument {
			c.AbortWithError(http.StatusBadRequest, err)
		} else if code == codes.NotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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
