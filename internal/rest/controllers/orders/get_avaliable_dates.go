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

func GetAvaliableDates(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := middleware.GetCustomer(c)
        if err != nil {
            c.AbortWithError(http.StatusUnauthorized, err)
            return
        }
		companyID := c.Param("company_id")
		if err := middleware.IsProperObjectIDHex(companyID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		serviceID := c.Param("service_id")
		if err := middleware.IsProperObjectIDHex(serviceID); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
            return
		}
        intValue, err := middleware.GetIntArg(c, "service_duration", 10, 480)
        if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
        }
        if intValue == nil {
			c.AbortWithError(
                http.StatusBadRequest, 
                errors.New("service duration field is required"),
            )
			return
        }
        serviceDuration := int32(*intValue)
        date, err := middleware.GetDateTime("date")
        if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
            return
        }
		var startValue *string = nil
		var query string = c.DefaultQuery("startValue", "")
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

		client := schedulerpb.NewApiClient(conns.GetSchedulerConn())
        message := schedulerpb.AvaliableTimeSlotsRequest{
            CompanyId: &companyID,
            ServiceId: &serviceID,
            ServiceDuration: &serviceDuration,
            Date: date,
            StartValue: startValue,
            NPerPage: &nPerPage,
        }
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
        reply, err := client.FindManyAvaliableTimeSlots(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
        c.JSON(http.StatusOK, reply)
	}
}
