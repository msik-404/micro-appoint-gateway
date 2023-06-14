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

func UpdateCustomer(c *gin.Context) {
	customer, err := middleware.GetCustomer(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	type CustomerPlain struct {
		Mail     *string `json:"mail" binding:"omitempty,max=30"`
		PlainPwd *string `json:"pwd" bidning:"omitempty,max=72"`
		Name     *string `json:"name" binding:"omitempty,max=30"`
		Surname  *string `json:"surname" binding:"omitempty,max=30"`
	}
	var customerUpdatePlain CustomerPlain
	if err := c.BindJSON(&customerUpdatePlain); err != nil {
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
