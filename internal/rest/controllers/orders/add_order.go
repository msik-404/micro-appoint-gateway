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

func AddOrder(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer, err := middleware.GetCustomer(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		newOrderPlain, err := middleware.GetData[middleware.Order](c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		companyID := newOrderPlain.CompanyID
		if err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        serviceID := newOrderPlain.ServiceID
		if err := middleware.IsProperObjectIDHex(serviceID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        employeeID := newOrderPlain.EmployeeID
		if err := middleware.IsProperObjectIDHex(employeeID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		startTime, err := middleware.GetDateTime(time.RFC3339, newOrderPlain.StartTime)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        endTime, err := middleware.GetDateTime(time.RFC3339, newOrderPlain.EndTime)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		client := schedulerpb.NewApiClient(conns.GetSchedulerConn())
        message := schedulerpb.AddOrderRequest{
            CustomerId: &customer.ID,
            CompanyId: &companyID,
            ServiceId: &serviceID,
            EmployeeId: &employeeID,
            StartTime: startTime,
            EndTime: endTime,
        }
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
        reply, err := client.AddOrder(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
        c.JSON(http.StatusOK, reply)
	}
}
