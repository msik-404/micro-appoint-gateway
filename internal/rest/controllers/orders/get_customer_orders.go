package orders

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/scheduler/schedulerpb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func GetCustomerOrders(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer, err := middleware.GetCustomer(c)
        if err != nil {
            c.AbortWithError(http.StatusUnauthorized, err)
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
            CustomerId: &customer.ID,
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
