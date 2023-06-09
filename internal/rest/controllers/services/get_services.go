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
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/communication"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func GetServices() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		companyID := c.Param("id")
		if _, err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var query string = c.DefaultQuery("startValue", "")
		var startValue *string = nil
		if query != "" {
			startValue = &query
			if _, err := middleware.IsProperObjectIDHex(*startValue); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
		nPerPage, err := middleware.GetNPerPageValue(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(companies.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := communication.NewApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		message := communication.ServicesRequest{
			CompanyId:  companyID,
			StartValue: startValue,
			NPerPage:   &nPerPage,
		}
		reply, err := client.FindManyServices(ctx, &message)
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
		c.JSON(http.StatusOK, reply.Services)
	}
	return gin.HandlerFunc(fn)
}
