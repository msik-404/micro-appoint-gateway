package grpc

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyGrpcClient[T any] struct {
    Client T
    Conn *grpc.ClientConn
}

func GetMyGrpcClient[T any](
    connString string, 
    getClient func(grpc.ClientConnInterface) T,
) (*MyGrpcClient[T], error) {
	var conn *grpc.ClientConn
    conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
        return nil, err
	}
	defer conn.Close()
	client := getClient(conn)
    return &MyGrpcClient[T]{Client: client, Conn: conn}, nil
}

func GrpcCodeToHttpCode(err error) int {
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
