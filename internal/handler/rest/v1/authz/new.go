package authz

import (
	"log/slog"
	"sync"

	"github.com/phillong97254/backoffice/config"
	"github.com/phillong97254/backoffice/internal/gateway"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz/account"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz/wallet"
	accountSvc "github.com/phillong97254/backoffice/internal/service/account"
	walletSvc "github.com/phillong97254/backoffice/internal/service/wallet"
)

type impl struct {
	cfg    *config.Config
	logger *slog.Logger
	gwy    gateway.Gateway

	account     account.Handler
	accountOnce sync.Once

	wallet     wallet.Handler
	walletOnce sync.Once
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Handler {
	return &impl{cfg: cfg, logger: logger, gwy: gwy}
}

func (i *impl) GetAccountHandler() account.Handler {
	i.accountOnce.Do(func() { i.account = account.New(i.logger, accountSvc.New(i.logger, i.gwy)) })
	return i.account
}

func (i *impl) GetWalletHandler() wallet.Handler {
	i.walletOnce.Do(func() { i.wallet = wallet.New(i.logger, walletSvc.New(i.cfg, i.logger, i.gwy)) })
	return i.wallet
}
