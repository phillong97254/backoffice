package account

import "net/http"

type Handler interface {
	// GET /api/me
	GetMe(w http.ResponseWriter, r *http.Request)
}
