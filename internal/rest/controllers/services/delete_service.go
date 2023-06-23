package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpctohttp"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func DeleteService(c *gin.Context) {
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
	serviceID := c.Param("service_id")
	if err := middleware.IsProperObjectIDHex(serviceID); err != nil {
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

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(companies.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	client := companiespb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := companiespb.DeleteServiceRequest{
		CompanyId: &companyID,
		Id:        &serviceID,
	}
	reply, err := client.DeleteService(ctx, &message)

	if err != nil {
        status := grpctohttp.GrpcCodeToHttpCode(err)
        c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
