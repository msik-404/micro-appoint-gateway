package services

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func UpdateService() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		serviceID := c.Param("id")
		if _, err := middleware.IsProperObjectIDHex(serviceID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var message companiespb.UpdateServiceRequest
		if err := c.BindJSON(&message); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		message.Id = serviceID
		if _, err := middleware.IsProperString(message.Name, 30); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperInteger(message.Price, 0, 1000000); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperInteger(message.Duration, 0, 480); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(message.Description, 300); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(companies.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := companiespb.NewApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.UpdateService(ctx, &message)
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
	return gin.HandlerFunc(fn)
}
