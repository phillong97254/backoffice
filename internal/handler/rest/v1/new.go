package v1

import (
	"log/slog"

	"github.com/phil/backoffice/config"
	"github.com/phil/backoffice/internal/gateway"
	"github.com/phil/backoffice/internal/handler/rest/v1/authz"
	"github.com/phil/backoffice/internal/handler/rest/v1/public"
)

type impl struct {
	public public.Handler
	authz  authz.Handler
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Handler {
	return &impl{
		public: public.New(logger, gwy),
		authz:  authz.New(cfg, logger, gwy),
	}
}

func (i *impl) GetPublicHandler() public.Handler { return i.public }
func (i *impl) GetAuthzHandler() authz.Handler   { return i.authz }
