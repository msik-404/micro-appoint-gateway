package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
)

func LoginCustomer(c *gin.Context) {
	type LoginPlain struct {
		Mail     string `json:"mail" binding:"required,max=30"`
		PlainPwd string `json:"pwd" bidning:"required,max=72"`
	}
	var loginPlain LoginPlain
	if err := c.BindJSON(&loginPlain); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := userspb.CustomerCredentialsRequest{
		Mail: &loginPlain.Mail,
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(users.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	client := userspb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.FindOneCustomerCredentials(ctx, &message)

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
	err = bcrypt.CompareHashAndPassword(
		[]byte(reply.GetHashedPwd()),
		[]byte(loginPlain.PlainPwd),
	)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	tokenStr, err := auth.CreateJWT(reply.GetId())
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
