package users 

import (
    mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
    "github.com/msik-404/micro-appoint-gateway/internal/grpc/users/userspb"
)

const (
	ConnString = "users:50051"
)

func GetClient() (*mygrpc.MyGrpcClient[userspb.ApiClient], error) {
    return mygrpc.GetMyGrpcClient[userspb.ApiClient](ConnString, userspb.NewApiClient)
}
