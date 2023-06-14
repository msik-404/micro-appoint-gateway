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
)

type Customer struct {
	ID      string
	Mail    string
	Name    string
	Surname string
}

func RequireCustomerAuth(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
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
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
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
            ID:      ownerID,
            Mail:    reply.GetMail(),
            Name:    reply.GetName(),
            Surname: reply.GetSurname(),
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
