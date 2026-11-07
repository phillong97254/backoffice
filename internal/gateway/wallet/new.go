package wallet

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/phil/backoffice/config"
	walletv1 "github.com/phil/backoffice/pkg/gen/wallet/v1"
	auth "github.com/phil/go-lib/authkit"
)

type impl struct {
	conn   *grpc.ClientConn
	client walletv1.WalletServiceClient
}

// New tạo gRPC client tới wallet-service, interceptor tự forward token của admin.
// Production: dùng TLS/mTLS thay cho insecure.
func New(cfg *config.Config) (Gateway, error) {
	conn, err := grpc.NewClient(cfg.WalletService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(auth.ForwardTokenInterceptor()))
	if err != nil {
		return nil, err
	}
	return &impl{conn: conn, client: walletv1.NewWalletServiceClient(conn)}, nil
}

func (i *impl) Close() error { return i.conn.Close() }
