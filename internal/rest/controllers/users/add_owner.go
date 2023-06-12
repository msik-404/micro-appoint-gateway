package users

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
)

func AddOwner(c *gin.Context) {
	type OwnerPlain struct {
		Mail     *string `json:"mail" binding:"required,max=30"`
		PlainPwd *string `json:"pwd" bidning:"required,max=72"`
		Name     *string `json:"name" binding:"omitempty,max=30"`
		Surname  *string `json:"surname" binding:"omitempty,max=30"`
	}
	var newOwnerPlain OwnerPlain
	if err := c.BindJSON(&newOwnerPlain); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hash, err := auth.HashAndSalt([]byte(*newOwnerPlain.PlainPwd))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hashedPwd := string(hash)
	message := userspb.AddOwnerRequest{
		Mail:      newOwnerPlain.Mail,
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
