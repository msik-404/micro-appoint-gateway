package employees 

import (
    mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
    "github.com/msik-404/micro-appoint-gateway/internal/grpc/employees/employeespb"
)

const (
	ConnString = "employees:50051"
)

func GetClient() (*mygrpc.MyGrpcClient[employeespb.ApiClient], error) {
    return mygrpc.GetMyGrpcClient[employeespb.ApiClient](ConnString, employeespb.NewApiClient)
}
