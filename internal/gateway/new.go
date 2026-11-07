package gateway

import (
	"fmt"

	"github.com/phil/backoffice/config"
	"github.com/phil/backoffice/internal/gateway/auth"
	"github.com/phil/backoffice/internal/gateway/user"
	"github.com/phil/backoffice/internal/gateway/wallet"
)

type impl struct {
	auth   auth.Gateway
	user   user.Gateway
	wallet wallet.Gateway
}

// New khởi tạo mọi client ngay lúc start (fail fast nếu cấu hình sai).
func New(cfg *config.Config) (Gateway, error) {
	walletGwy, err := wallet.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("init wallet gateway: %w", err)
	}
	return &impl{
		auth:   auth.New(cfg),
		user:   user.New(cfg),
		wallet: walletGwy,
	}, nil
}

func (i *impl) Auth() auth.Gateway     { return i.auth }
func (i *impl) User() user.Gateway     { return i.user }
func (i *impl) Wallet() wallet.Gateway { return i.wallet }
func (i *impl) Close() error           { return i.wallet.Close() }
