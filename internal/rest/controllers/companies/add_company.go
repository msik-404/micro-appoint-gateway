package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func AddCompany(c *gin.Context) {
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

    myCompaniesClient, err := companies.GetClient()
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    defer myCompaniesClient.Conn.Close()
    companiesClient := myCompaniesClient.Client
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
		status := mygrpc.GrpcCodeToHttpCode(err)
		c.AbortWithError(status, err)
		return
	}

    myUsersClient, err := users.GetClient()
	// If connection could not be established, roll back changes.
	if err != nil {
		deleteCompanyMessage := companiespb.DeleteCompanyRequest{
			Id: addCompanyReply.Id,
		}
		_, deleteErr := companiesClient.DeleteCompany(ctx, &deleteCompanyMessage)
		if deleteErr != nil {
			err = deleteErr
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer myUsersClient.Conn.Close()
	usersClient := myUsersClient.Client
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
		status := mygrpc.GrpcCodeToHttpCode(err)
		c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, usersReply)
}
