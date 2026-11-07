package routes

import (
	"net/http"

	v1 "github.com/phil/backoffice/internal/handler/rest/v1"
)

func v1PublicRoutes(mux *http.ServeMux, h v1.Handler) {
	authHandler := h.GetPublicHandler().GetAuthHandler()
	mux.HandleFunc("POST /api/login", authHandler.Login)
}
