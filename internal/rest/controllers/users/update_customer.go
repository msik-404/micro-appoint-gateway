package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpctohttp"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func UpdateCustomer(c *gin.Context) {
	customer, err := middleware.GetCustomer(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

    customerUpdatePlain, err := middleware.GetData[middleware.UserUpdate](c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var newHashedPwd *string
	if customerUpdatePlain.PlainPwd != nil {
		hash, err := auth.HashAndSalt([]byte(*customerUpdatePlain.PlainPwd))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		hashedPwd := string(hash)
		newHashedPwd = &hashedPwd
	}
	message := userspb.UpdateCustomerRequest{
		Id:        &customer.ID,
		Mail:      customerUpdatePlain.Mail,
		HashedPwd: newHashedPwd,
		Name:      customerUpdatePlain.Name,
		Surname:   customerUpdatePlain.Surname,
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
	reply, err := client.UpdateCustomer(ctx, &message)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
