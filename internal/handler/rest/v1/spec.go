package v1

import (
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/admin"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/public"
)

type Handler interface {
	GetPublicHandler() public.Handler
	GetAuthzHandler() authz.Handler
	GetAdminHandler() admin.Handler
}
