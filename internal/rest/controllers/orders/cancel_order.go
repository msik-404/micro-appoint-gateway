package orders

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/msik-404/micro-appoint-gateway/internal/rabbit"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

type Request struct {
	CustomerID string `json:"customer_id"`
	OrderID    string `json:"order_id"`
}

func CancelOrder(conns *rabbit.RabbitConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer, err := middleware.GetCustomer(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		cancelRequest, err := middleware.GetData[middleware.CancelRequest](c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ch, err := conns.Conn.Channel()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		requestQueue, err := ch.QueueDeclare(
			"request-cancel", // name
			false,            // durable
			false,            // delete when unused
			false,            // exclusive
			false,            // no-wait
			nil,              // arguments
		)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		request := Request{
			CustomerID: customer.ID,
			OrderID:    cancelRequest.OrderID,
		}
        data, err := json.Marshal(request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
        ch.Publish(
			"",
			requestQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:  "application/json",
				Body:         data,
				DeliveryMode: amqp.Persistent,
			})

		c.JSON(http.StatusOK, gin.H{})
	}
}
