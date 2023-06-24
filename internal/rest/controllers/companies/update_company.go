package companies

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func UpdateCompany(c *gin.Context) {
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

	companyUpdate, err := middleware.GetData[middleware.Company](c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := companiespb.UpdateCompanyRequest{
		Id:               &companyID,
		Name:             companyUpdate.Name,
		Type:             companyUpdate.Type,
		Localisation:     companyUpdate.Localisation,
		ShortDescription: companyUpdate.ShortDescription,
		LongDescription:  companyUpdate.LongDescription,
	}

    myClient, err := companies.GetClient()
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    defer myClient.Conn.Close()
    client := myClient.Client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.UpdateCompany(ctx, &message)

	if err != nil {
		status := mygrpc.GrpcCodeToHttpCode(err)
		c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply)
}
