package role

import (
	"log/slog"

	roleSvc "github.com/phillong97254/backoffice/internal/service/role"
)

type impl struct {
	logger  *slog.Logger
	roleSvc roleSvc.Service
}

func New(logger *slog.Logger, svc roleSvc.Service) Handler {
	return &impl{logger: logger.With("handler", "role"), roleSvc: svc}
}
