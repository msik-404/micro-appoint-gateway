package companies

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	type Company struct {
		Name             *string `json:"name" binding:"omitempty,max=30"`
		Type             *string `json:"type" binding:"omitempty,max=30"`
		Localisation     *string `json:"localisation" binding:"omitempty,max=60"`
		ShortDescription *string `json:"short_description" binding:"omitempty,max=150"`
		LongDescription  *string `json:"long_description" binding:"omitempty,max=300"`
	}
	var newCompany Company
	if err := c.BindJSON(&newCompany); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := companiespb.AddCompanyRequest{
		Name:             newCompany.Name,
		Type:             newCompany.Type,
		Localisation:     newCompany.Localisation,
		ShortDescription: newCompany.ShortDescription,
		LongDescription:  newCompany.LongDescription,
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(companies.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	companiesClient := companiespb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	companiesReply, err := companiesClient.AddCompany(ctx, &message)

	if err != nil {
		code := status.Code(err)
		if code == codes.InvalidArgument {
			c.AbortWithError(http.StatusBadRequest, err)
		} else if code == codes.NotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else if code == codes.AlreadyExists {
            c.AbortWithError(http.StatusConflict, err)
        } else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

    conn.Close()
	conn, err = grpc.Dial(users.ConnString, grpc.WithInsecure())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
    defer conn.Close()
	usersMessage := userspb.AddOwnedCompanyRequest{
		Id:        &owner.ID,
		CompanyId: companiesReply.Id,
	}
	usersClient := userspb.NewApiClient(conn)
	usersReply, err := usersClient.AddOwnedCompany(ctx, &usersMessage)

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
	c.JSON(http.StatusOK, usersReply)
}
