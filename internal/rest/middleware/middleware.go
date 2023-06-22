package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
	"github.com/msik-404/micro-appoint-gateway/internal/strtime"
)

type Company struct {
	Name             *string `json:"name" binding:"omitempty,max=30"`
	Type             *string `json:"type" binding:"omitempty,max=30"`
	Localisation     *string `json:"localisation" binding:"omitempty,max=60"`
	ShortDescription *string `json:"short_description" binding:"omitempty,max=150"`
	LongDescription  *string `json:"long_description" binding:"omitempty,max=300"`
}

type Service struct {
	Name        *string `json:"name" binding:"omitempty,max=30"`
	Price       *int32  `json:"price" binding:"omitempty,min=0,max=1000000"`
	Duration    *int32  `json:"duration" binding:"omitempty,min=0,max=480"`
	Description *string `json:"description" binding:"omitempty,max=300"`
}

type Employee struct {
	Name       *string               `json:"name" binding:"omitempty,max=30"`
	Surname    *string               `json:"surname" binding:"omitempty,max=30"`
	WorkTimes  *strtime.WorkTimesStr `json:"work_times" binding:"omitempty"`
	Competence []string              `json:"competence" binding:"omitempty"`
}

type User struct {
	Mail     string  `json:"mail" binding:"required,max=30"`
	PlainPwd string  `json:"pwd" bidning:"required,max=72"`
	Name     *string `json:"name" binding:"omitempty,max=30"`
	Surname  *string `json:"surname" binding:"omitempty,max=30"`
}

type UserUpdate struct {
	Mail     *string `json:"mail" binding:"omitempty,max=30"`
	PlainPwd *string `json:"pwd" bidning:"omitempty,max=72"`
	Name     *string `json:"name" binding:"omitempty,max=30"`
	Surname  *string `json:"surname" binding:"omitempty,max=30"`
}

func bind[T any](c *gin.Context, key string) {
	var generic T
	if err := c.BindJSON(&generic); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Set(key, generic)
}

func BindAuth[T any](c *gin.Context) {
	bind[T](c, "auth")
}

func Bind[T any](c *gin.Context) {
	bind[T](c, "data")
}

func BindWithAuth[T any](c *gin.Context) {
	type AuthGeneric struct {
		Auth auth.Token `json:"auth" binding:"required"`
		Data T          `json:"data" binding:"required"`
	}
	var authGeneric AuthGeneric
	if err := c.BindJSON(&authGeneric); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Set("auth", authGeneric.Auth)
	c.Set("data", authGeneric.Data)
}

func GetData[T any](c *gin.Context) (*T, error) {
	result, ok := c.Get("data")
	if !ok {
		return nil, errors.New("data field not set")
	}
	data := result.(T)
	return &data, nil
}

func GetAuth(c *gin.Context) (*auth.Token, error) {
	result, ok := c.Get("auth")
	if !ok {
		return nil, errors.New("user unauthorized")
	}
	data := result.(auth.Token)
	return &data, nil
}

type AuthHeader struct {
    Auth string `header:"Authentication" binding:"required"`
}

func getAuthHeader(c *gin.Context) (*string, error) {
    var header AuthHeader
    if err := c.BindHeader(&header); err != nil {
        return nil, err
    }
    return &header.Auth, nil
}

type Customer struct {
	ID      string
	Mail    string
	Name    string
	Surname string
}

func RequireCustomerAuth(c *gin.Context) {
	tokenStr, err := getAuthHeader(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := jwt.Parse(*tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// check if that user is present in database
		customerID := claims["user_id"].(string)
		message := userspb.CustomerRequest{
			Id: &customerID,
		}

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(users.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer conn.Close()
		client := userspb.NewApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.FindOneCustomer(ctx, &message)
		if err != nil {
			code := status.Code(err)
			if code == codes.InvalidArgument {
				c.AbortWithError(http.StatusBadRequest, err)
			} else if code == codes.NotFound {
				c.AbortWithError(http.StatusUnauthorized, err)
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}
		customer := Customer{
			ID:      customerID,
			Mail:    reply.GetMail(),
			Name:    reply.GetName(),
			Surname: reply.GetSurname(),
		}
		c.Set("customer", customer)
		c.Next()
		return
	}
	c.AbortWithError(http.StatusUnauthorized, err)
	return
}

func GetCustomer(c *gin.Context) (Customer, error) {
	result, ok := c.Get("customer")
	if !ok {
		return Customer{}, errors.New("Customer user unauthorized")
	}
	customer := result.(Customer)
	return customer, nil
}

type Owner struct {
	ID        string
	Mail      string
	Name      string
	Surname   string
	Companies map[string]struct{}
}

func RequireOwnerAuth(c *gin.Context) {
	tokenStr, err := getAuthHeader(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := jwt.Parse(*tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// check if that user is present in database
		ownerID := claims["user_id"].(string)
		message := userspb.OwnerRequest{
			Id: &ownerID,
		}

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(users.ConnString, grpc.WithInsecure())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer conn.Close()
		client := userspb.NewApiClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reply, err := client.FindOneOwner(ctx, &message)

		if err != nil {
			code := status.Code(err)
			if code == codes.InvalidArgument {
				c.AbortWithError(http.StatusBadRequest, err)
			} else if code == codes.NotFound {
				c.AbortWithError(http.StatusUnauthorized, err)
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}
		owner := Owner{
			ID:        ownerID,
			Mail:      reply.GetMail(),
			Name:      reply.GetName(),
			Surname:   reply.GetSurname(),
			Companies: make(map[string]struct{}),
		}
		for _, company := range reply.GetCompanies() {
			owner.Companies[company] = struct{}{}
		}
		c.Set("owner", owner)
		c.Next()
		return
	}
	c.AbortWithError(http.StatusUnauthorized, err)
	return
}

func GetOwner(c *gin.Context) (Owner, error) {
	result, ok := c.Get("owner")
	if !ok {
		return Owner{}, errors.New("Owner user unauthorized")
	}
	owner := result.(Owner)
	return owner, nil
}

func IsProperObjectIDHex(hex string) error {
	if len(hex) != 24 {
		return errors.New("This is not proper hex value for objectID")
	}
	return nil
}

func GetNPerPageValue(c *gin.Context) (int64, error) {
	query := c.DefaultQuery("nPerPage", "30")
	nPerPage, err := strconv.Atoi(query)
	if err != nil {
		return int64(nPerPage), err
	}
	if nPerPage < 0 {
		return int64(nPerPage), errors.New("nPerPage should be positive number")
	}
	return int64(nPerPage), nil
}
