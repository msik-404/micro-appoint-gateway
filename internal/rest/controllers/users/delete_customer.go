package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func DeleteCustomer(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer, err := middleware.GetCustomer(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		client := userspb.NewApiClient(conns.GetUsersConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		message := userspb.DeleteCustomerRequest{Id: &customer.ID}
		reply, err := client.DeleteCustomer(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
		c.JSON(http.StatusOK, reply)
	}
}
