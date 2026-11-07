package role

import (
	"log/slog"
	"time"

	"github.com/phillong97254/backoffice/config"
	"github.com/phillong97254/backoffice/internal/gateway"
)

type impl struct {
	logger  *slog.Logger
	gwy     gateway.Gateway
	timeout time.Duration
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Service {
	return &impl{
		logger:  logger.With("service", "role"),
		gwy:     gwy,
		timeout: time.Duration(cfg.AuthzService.Timeout) * time.Second,
	}
}
