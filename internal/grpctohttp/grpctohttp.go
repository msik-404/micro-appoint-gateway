package grpctohttp

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcCodeToHttpCode(err error) int {
	switch code := status.Code(err); code {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
    case codes.Unauthenticated :
        return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
