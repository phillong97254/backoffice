package wallet

import (
	"log/slog"

	walletSvc "github.com/phillong97254/backoffice/internal/service/wallet"
)

type impl struct {
	logger    *slog.Logger
	walletSvc walletSvc.Service
}

func New(logger *slog.Logger, svc walletSvc.Service) Handler {
	return &impl{logger: logger.With("handler", "wallet"), walletSvc: svc}
}
