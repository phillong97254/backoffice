package authz

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	auth "github.com/phillong97254/go-lib/authkit"

	"github.com/phillong97254/backoffice/config"
	authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"
)

type impl struct {
	conn   *grpc.ClientConn
	client authzv1.AuthzServiceClient
}

// New tạo gRPC client tới authz-service, interceptor tự forward token của admin
// (authz-service cần biết ai đang cấp/thu quyền để kiểm tra và ghi audit).
// Production: dùng TLS/mTLS thay cho insecure.
func New(cfg *config.Config) (Gateway, error) {
	conn, err := grpc.NewClient(cfg.AuthzService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(auth.ForwardTokenInterceptor()))
	if err != nil {
		return nil, err
	}
	return &impl{conn: conn, client: authzv1.NewAuthzServiceClient(conn)}, nil
}

func (i *impl) Close() error { return i.conn.Close() }
