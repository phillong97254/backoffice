package account

import (
	"log/slog"

	"github.com/phillong97254/backoffice/internal/gateway"
)

type impl struct {
	logger *slog.Logger
	gwy    gateway.Gateway
}

func New(logger *slog.Logger, gwy gateway.Gateway) Service {
	return &impl{logger: logger.With("service", "account"), gwy: gwy}
}
