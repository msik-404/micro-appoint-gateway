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
		code := status.Code(err)
		if code == codes.InvalidArgument {
			c.AbortWithError(http.StatusBadRequest, err)
		} else if code == codes.NotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else if code == codes.AlreadyExists {
            c.AbortWithError(http.StatusConflict, err)
        } else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, reply)
}
