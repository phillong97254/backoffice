package gateway

import (
	"github.com/phil/backoffice/internal/gateway/auth"
	"github.com/phil/backoffice/internal/gateway/user"
	"github.com/phil/backoffice/internal/gateway/wallet"
)

// Gateway gom các client gọi sang service khác.
type Gateway interface {
	Auth() auth.Gateway     // auth-service (REST, OAuth2)
	User() user.Gateway     // user-service (REST, token client_credentials của gateway)
	Wallet() wallet.Gateway // wallet-service (gRPC, forward token của admin)
	Close() error
}
