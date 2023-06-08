package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/communication"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func UpdateCompany() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		companyID := c.Param("id")
		if _, err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var message communication.UpdateCompanyRequest
		if err := c.BindJSON(&message); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		message.Id = companyID
		if _, err := middleware.IsProperString(message.Name, 30); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(message.Type, 30); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(message.Localisation, 60); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(message.ShortDescription, 150); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, err := middleware.IsProperString(message.LongDescription, 300); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(companies.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := communication.NewApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.UpdateCompany(ctx, &message)
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
