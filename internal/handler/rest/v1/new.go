package v1

import (
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/admin"
	"log/slog"

	"github.com/phillong97254/backoffice/config"
	"github.com/phillong97254/backoffice/internal/gateway"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/public"
)

type impl struct {
	public public.Handler
	authz  authz.Handler
	admin  admin.Handler
}

func New(cfg *config.Config, logger *slog.Logger, gwy gateway.Gateway) Handler {
	return &impl{
		public: public.New(logger, gwy),
		authz:  authz.New(cfg, logger, gwy),
		admin:  admin.New(cfg, logger, gwy),
	}
}

func (i *impl) GetPublicHandler() public.Handler { return i.public }
func (i *impl) GetAuthzHandler() authz.Handler   { return i.authz }
func (i *impl) GetAdminHandler() admin.Handler   { return i.admin }
