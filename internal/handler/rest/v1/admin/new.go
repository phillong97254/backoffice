package admin

import (
	"log/slog"
	"sync"

	"github.com/phillong97254/backoffice/config"
	"github.com/phillong97254/backoffice/internal/gateway"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/admin/role"
	roleSvc "github.com/phillong97254/backoffice/internal/service/role"
)

type impl struct {
	cfg    *config.Config
	logger *slog.Logger
	gwy    gateway.Gateway

	role     role.Handler
	roleOnce sync.Once
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Handler {
	return &impl{cfg: cfg, logger: logger, gwy: gwy}
}

func (i *impl) GetRoleHandler() role.Handler {
	i.roleOnce.Do(func() { i.role = role.New(i.logger, roleSvc.New(i.cfg, i.logger, i.gwy)) })
	return i.role
}
