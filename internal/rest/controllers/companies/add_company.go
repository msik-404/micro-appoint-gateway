package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func AddCompany(conns *mygrpc.GRPCConns) gin.HandlerFunc {
	return func(c *gin.Context) {
		owner, err := middleware.GetOwner(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		newCompany, err := middleware.GetData[middleware.Company](c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		companiesClient := companiespb.NewApiClient(conns.GetCompaniesConn())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		addCompanyMessage := companiespb.AddCompanyRequest{
			Name:             newCompany.Name,
			Type:             newCompany.Type,
			Localisation:     newCompany.Localisation,
			ShortDescription: newCompany.ShortDescription,
			LongDescription:  newCompany.LongDescription,
		}
		addCompanyReply, err := companiesClient.AddCompany(ctx, &addCompanyMessage)

		if err != nil {
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}

		usersClient := userspb.NewApiClient(conns.GetUsersConn())
		usersMessage := userspb.AddOwnedCompanyRequest{
			Id:        &owner.ID,
			CompanyId: addCompanyReply.Id,
		}
		usersReply, err := usersClient.AddOwnedCompany(ctx, &usersMessage)

		// If operation cloud not be performed, roll back changes.
		if err != nil {
			deleteCompanyMessage := companiespb.DeleteCompanyRequest{
				Id: addCompanyReply.Id,
			}
			_, deleteErr := companiesClient.DeleteCompany(ctx, &deleteCompanyMessage)
			if deleteErr != nil {
				err = deleteErr
			}
			status := mygrpc.GRPCCodeToHTTPCode(err)
			c.AbortWithError(status, err)
			return
		}
		c.JSON(http.StatusOK, usersReply)
	}
}
