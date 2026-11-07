package auth

import (
	"log/slog"

	authSvc "github.com/phil/backoffice/internal/service/auth"
)

type impl struct {
	logger  *slog.Logger
	authSvc authSvc.Service
}

func New(logger *slog.Logger, svc authSvc.Service) Handler {
	return &impl{logger: logger.With("handler", "auth"), authSvc: svc}
}
