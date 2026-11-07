package account

import (
	"log/slog"

	accountSvc "github.com/phil/backoffice/internal/service/account"
)

type impl struct {
	logger     *slog.Logger
	accountSvc accountSvc.Service
}

func New(logger *slog.Logger, svc accountSvc.Service) Handler {
	return &impl{logger: logger.With("handler", "account"), accountSvc: svc}
}
