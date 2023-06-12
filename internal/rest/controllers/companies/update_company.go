package companies

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	type Company struct {
		Name             *string `json:"name" binding:"omitempty,max=30"`
		Type             *string `json:"type" binding:"omitempty,max=30"`
		Localisation     *string `json:"localisation" binding:"omitempty,max=60"`
		ShortDescription *string `json:"short_description" binding:"omitempty,max=150"`
		LongDescription  *string `json:"long_description" binding:"omitempty,max=300"`
	}
	var companyUpdate Company
	if err := c.BindJSON(&companyUpdate); err != nil {
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
	reply, err := client.UpdateCompany(ctx, &message)

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
	c.JSON(http.StatusOK, reply)
}
