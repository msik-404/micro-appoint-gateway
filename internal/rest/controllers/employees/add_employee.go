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

func AddEmployee() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		companyID := c.Param("id")
		if _, err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		type Employee struct {
			Name       *string               `json:"name,omitempty"`
			Surname    *string               `json:"surname,omitempty"`
			WorkTimes  *strtime.WorkTimesStr `json:"work_times,omitempty"`
			Competence []string              `json:"competence,omitempty"`
		}
		var employeeNew Employee

		if err := c.BindJSON(&employeeNew); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(employeeNew.Name, 30); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(employeeNew.Surname, 30); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        message := employeespb.AddEmployeeRequest{
            CompanyId: companyID,
            Name:      employeeNew.Name,
            Surname:   employeeNew.Surname,
        }
		if employeeNew.WorkTimes != nil {
			workTimes, err := strtime.ToWorkTimes(employeeNew.WorkTimes)
            if err != nil {
                c.AbortWithError(http.StatusBadRequest, err)
                return
            }
            message.WorkTimes = workTimes
		}
		for _, hex := range employeeNew.Competence {
			if _, err := middleware.IsProperObjectIDHex(hex); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			message.Competence = append(message.Competence, hex)
		}

		var conn *grpc.ClientConn
        conn, err := grpc.Dial(employees.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := employeespb.NewApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.AddEmployee(ctx, &message)
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
		c.JSON(http.StatusOK, reply)
	}
	return gin.HandlerFunc(fn)
}
