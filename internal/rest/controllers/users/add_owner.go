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

func AddOwner(c *gin.Context) {
    newOwnerPlain, err := middleware.GetData[middleware.User](c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hash, err := auth.HashAndSalt([]byte(newOwnerPlain.PlainPwd))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hashedPwd := string(hash)
	message := userspb.AddOwnerRequest{
		Mail:      &newOwnerPlain.Mail,
		HashedPwd: &hashedPwd,
		Name:      newOwnerPlain.Name,
		Surname:   newOwnerPlain.Surname,
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
	reply, err := client.AddOwner(ctx, &message)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
