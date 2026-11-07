// Package grpcerr đổi lỗi gRPC từ các service phía sau (wallet-service, authz-service)
// sang lỗi HTTP (apperr) cho frontend. Dùng ở tầng service.
package grpcerr

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/phillong97254/go-lib/apperr"
)

var grpcToHTTP = map[codes.Code]int{
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.NotFound:           http.StatusNotFound,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.FailedPrecondition: http.StatusConflict,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.Unavailable:        http.StatusServiceUnavailable,
}

func FromGRPC(err error) error {
	st := status.Convert(err)
	code, ok := grpcToHTTP[st.Code()]
	if !ok {
		return &apperr.Error{Status: http.StatusInternalServerError, Message: st.Message(), Err: err}
	}
	return apperr.New(code, st.Message())
}
