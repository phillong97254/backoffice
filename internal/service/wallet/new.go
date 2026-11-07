package wallet

import (
	"log/slog"
	"time"

	"github.com/phil/backoffice/config"
	"github.com/phil/backoffice/internal/gateway"
)

type impl struct {
	logger  *slog.Logger
	gwy     gateway.Gateway
	timeout time.Duration
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Service {
	return &impl{
		logger:  logger.With("service", "wallet"),
		gwy:     gwy,
		timeout: time.Duration(cfg.WalletService.Timeout) * time.Second,
	}
}
