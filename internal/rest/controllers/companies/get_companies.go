package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func GetCompanies(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query string = c.DefaultQuery("startValue", "")
		var startValue *string = nil
		if query != "" {
			startValue = &query
			if err := middleware.IsProperObjectIDHex(*startValue); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
		nPerPage, err := middleware.GetNPerPageValue(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		client := companiespb.NewApiClient(conns.GetCompaniesConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		message := companiespb.CompaniesRequest{
			StartValue: startValue,
			NPerPage:   &nPerPage,
		}
		reply, err := client.FindManyCompanies(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
		c.JSON(http.StatusOK, reply.Companies)
	}
}
