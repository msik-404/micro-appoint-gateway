package orders

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/scheduler/schedulerpb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

// todo: make it so it returns strings instead of ids
func GetCompanyOrders(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		owner, err := middleware.GetOwner(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		companyID := c.Param("company_id")
		if err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if _, ok := owner.Companies[companyID]; !ok {
			c.AbortWithError(
				http.StatusUnauthorized,
				errors.New("This owner does not own this company"),
			)
			return
		}
        isCanceled, err := middleware.GetBoolArg(c, "is_canceled")
        if err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        nPerPage, err := middleware.GetNPerPageValue(c)
        if err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
		var query string = c.DefaultQuery("startValue", "")
		var startValue *string = nil
		if query != "" {
			startValue = &query
			if err := middleware.IsProperObjectIDHex(*startValue); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
        var startDate *int64 = nil
		query = c.DefaultQuery("startDate", "")
		if query != "" {
            startDate, err = middleware.GetDateTime(query)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
		client := schedulerpb.NewApiClient(conns.GetSchedulerConn())
		message := schedulerpb.OrdersRequest{
			CompanyId: &companyID,
            IsCanceled: isCanceled,
            NPerPage: &nPerPage,
            StartValue: startValue,
            StartDate: startDate,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
        reply, err := client.FindManyOrders(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
        c.JSON(http.StatusOK, reply.GetOrders())
	}
}
