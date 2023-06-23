package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpctohttp"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func DeleteCustomer(c *gin.Context) {
	customer, err := middleware.GetCustomer(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(users.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	client := userspb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := userspb.DeleteCustomerRequest{Id: &customer.ID}
	reply, err := client.DeleteCustomer(ctx, &message)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
