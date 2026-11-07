package v1

import (
	"github.com/phil/backoffice/internal/handler/rest/v1/authz"
	"github.com/phil/backoffice/internal/handler/rest/v1/public"
)

type Handler interface {
	GetPublicHandler() public.Handler
	GetAuthzHandler() authz.Handler
}
