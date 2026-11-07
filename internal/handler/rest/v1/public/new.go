package public

import (
	"log/slog"
	"sync"

	"github.com/phil/backoffice/internal/gateway"
	"github.com/phil/backoffice/internal/handler/rest/v1/public/auth"
	authSvc "github.com/phil/backoffice/internal/service/auth"
)

type impl struct {
	logger *slog.Logger
	gwy    gateway.Gateway

	auth     auth.Handler
	authOnce sync.Once
}

func New(logger *slog.Logger, gwy gateway.Gateway) Handler {
	return &impl{logger: logger, gwy: gwy}
}

func (i *impl) GetAuthHandler() auth.Handler {
	i.authOnce.Do(func() { i.auth = auth.New(i.logger, authSvc.New(i.logger, i.gwy)) })
	return i.auth
}
