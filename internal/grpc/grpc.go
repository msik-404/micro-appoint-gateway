package grpc

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/grpc/users"
)

type GRPCConns struct {
	CompaniesConn *grpc.ClientConn
	EmployeeSConn *grpc.ClientConn
	UsersConn     *grpc.ClientConn
}

func New() (*GRPCConns, error) {
	conns := GRPCConns{}
	companiesConn, err := grpc.Dial(companies.ConnString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conns.CompaniesConn = companiesConn
	employeesConn, err := grpc.Dial(employees.ConnString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conns.EmployeeSConn = employeesConn
	usersConn, err := grpc.Dial(users.ConnString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conns.UsersConn = usersConn
	return &conns, nil
}

func (conns *GRPCConns) GetCompaniesConn() *grpc.ClientConn {
	return conns.CompaniesConn
}

func (conns *GRPCConns) GetEmployeesConn() *grpc.ClientConn {
	return conns.EmployeeSConn
}

func (conns *GRPCConns) GetUsersConn() *grpc.ClientConn {
	return conns.UsersConn
}

func GRPCCodeToHTTPCode(err error) int {
	switch code := status.Code(err); code {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
