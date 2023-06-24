package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func LoginCustomer(c *gin.Context) {
	loginPlain, err := middleware.GetData[middleware.User](c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := userspb.CustomerCredentialsRequest{
		Mail: &loginPlain.Mail,
	}

    myClient, err := users.GetClient()
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    defer myClient.Conn.Close()
    client := myClient.Client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.FindOneCustomerCredentials(ctx, &message)

	if err != nil {
		status := mygrpc.GrpcCodeToHttpCode(err)
		if status == http.StatusNotFound {
			c.AbortWithError(http.StatusUnauthorized, err)
		} else {
			c.AbortWithError(status, err)
		}
		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(reply.GetHashedPwd()),
		[]byte(loginPlain.PlainPwd),
	)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := auth.CreateJWT(reply.GetId())
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
