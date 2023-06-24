package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func AddOwner(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		client := userspb.NewApiClient(conns.GetUsersConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.AddOwner(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
		c.JSON(http.StatusOK, reply)
	}
}
