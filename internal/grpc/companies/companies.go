package companies

import (
    mygrpc "github.com/msik-404/micro-appoint-gateway/internal/grpc"
    "github.com/msik-404/micro-appoint-gateway/internal/grpc/companies/companiespb"
)

const (
	ConnString = "companies:50051"
)

func GetClient() (*mygrpc.MyGrpcClient[companiespb.ApiClient], error) {
    return mygrpc.GetMyGrpcClient[companiespb.ApiClient](ConnString, companiespb.NewApiClient)
}
