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

func UpdateCustomer(conns *mygrpc.GRPCConns) gin.HandlerFunc {
return func (c *gin.Context) {
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

    client := userspb.NewApiClient(conns.GetUsersConn())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.UpdateCustomer(ctx, &message)

	if err != nil {
		status := mygrpc.GRPCCodeToHTTPCode(err)
		c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
}
