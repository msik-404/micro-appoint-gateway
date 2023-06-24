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

func GetCompaniesByIds(c *gin.Context) {
	owner, err := middleware.GetOwner(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	var query string = c.DefaultQuery("startValue", "")
	var startValue *string = nil
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
	if len(owner.Companies) == 0 {
		c.AbortWithError(
			http.StatusBadRequest,
			errors.New("You don't own any company"),
		)
		return
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
	companiesIDS := make([]string, len(owner.Companies))
	i := 0
	for k := range owner.Companies {
		companiesIDS[i] = k
		i++
	}
	message := companiespb.CompaniesByIdsRequest{
		Ids:        companiesIDS,
		StartValue: startValue,
		NPerPage:   &nPerPage,
	}
	reply, err := client.FindManyCompaniesByIds(ctx, &message)
	if err != nil {
		status := mygrpc.GrpcCodeToHttpCode(err)
		c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusOK, reply.Companies)
}
