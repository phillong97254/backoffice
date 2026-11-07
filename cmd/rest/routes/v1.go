package routes

import (
	"net/http"

	v1 "github.com/phil/backoffice/internal/handler/rest/v1"
	auth "github.com/phil/backoffice/pkg/authkit"
)

func V1Routes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	v1PublicRoutes(mux, h)
	v1AuthzRoutes(mux, verifier, h)
}
