package routes

import (
	"net/http"

	v1 "github.com/phil/backoffice/internal/handler/rest/v1"
	auth "github.com/phil/backoffice/pkg/authkit"
)

// Mọi /api/* (trừ login) cần JWT. API ví còn cần role admin (authz lớp 1, chặn sớm ở backoffice).
func v1AuthzRoutes(mux *http.ServeMux, verifier *auth.Verifier, h v1.Handler) {
	api := http.NewServeMux()
	admin := auth.Policy{Roles: []string{"admin"}}

	accountHandler := h.GetAuthzHandler().GetAccountHandler()
	api.HandleFunc("GET /api/me", accountHandler.GetMe)

	walletHandler := h.GetAuthzHandler().GetWalletHandler()
	api.HandleFunc("GET /api/wallets/{adminID}", auth.Require(admin, walletHandler.GetWallet))
	api.HandleFunc("POST /api/wallets/{adminID}/adjust", auth.Require(admin, walletHandler.AdjustBalance))

	mux.Handle("/api/", auth.Authenticate(verifier)(api))
}
