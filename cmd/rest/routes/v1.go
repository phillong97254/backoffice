package routes

import (
	"net/http"

	v1 "github.com/phillong97254/backoffice/internal/handler/rest/v1"
	auth "github.com/phillong97254/go-lib/authkit"
)

func V1Routes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	v1PublicRoutes(mux, h)
	v1AuthzRoutes(mux, verifier, h)
}
