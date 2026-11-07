package account

import (
	"log/slog"

	"github.com/phil/backoffice/internal/gateway"
)

type impl struct {
	logger *slog.Logger
	gwy    gateway.Gateway
}

func New(logger *slog.Logger, gwy gateway.Gateway) Service {
	return &impl{logger: logger.With("service", "account"), gwy: gwy}
}
