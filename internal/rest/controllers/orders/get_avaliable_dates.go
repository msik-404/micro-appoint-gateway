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
		query := c.DefaultQuery("date", "")
		date, err := middleware.GetDateTime(time.DateOnly, query)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var startValue *string = nil
		query = c.DefaultQuery("startValue", "")
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
			CompanyId:       &companyID,
			ServiceId:       &serviceID,
			ServiceDuration: &serviceDuration,
			Date:            date,
			StartValue:      startValue,
			NPerPage:        &nPerPage,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.FindManyAvaliableTimeSlots(ctx, &message)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
        var response []EmployeeTimeSlots
        for i := range reply.EmployeeTimeSlots {
            employeeTImeSlots := EmployeeTimeSlots{
                Id: reply.EmployeeTimeSlots[i].GetId(),
                Name: reply.EmployeeTimeSlots[i].GetName(),
                Surname: reply.EmployeeTimeSlots[i].GetSurname(),
            }
            for j := range reply.EmployeeTimeSlots[i].TimeSlots {
                startTime := time.Unix(reply.EmployeeTimeSlots[i].TimeSlots[j].GetStartTime(), 0)
                endTime := time.Unix(reply.EmployeeTimeSlots[i].TimeSlots[j].GetEndTime(), 0)
                timeSlot := TimeSlot{
                    StartTime: startTime.Format(time.RFC3339),
                    EndTime: endTime.Format(time.RFC3339),
                }
                employeeTImeSlots.TimeSlots = append(employeeTImeSlots.TimeSlots, timeSlot)
            }
            response = append(response, employeeTImeSlots)
        }
		c.JSON(http.StatusOK, response)
	}
}

type TimeSlot struct {
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

type EmployeeTimeSlots struct {
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Surname   string     `json:"surname,omitempty"`
	TimeSlots []TimeSlot `json:"time_slots,omitempty"`
}
