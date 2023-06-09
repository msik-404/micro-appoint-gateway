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
)

func DeleteEmployee() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		employeeID := c.Param("id")
		if _, err := middleware.IsProperObjectIDHex(employeeID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
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

        message := employeespb.DeleteEmployeeRequest{Id: employeeID}
		reply, err := client.DeleteEmployee(ctx, &message)
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
