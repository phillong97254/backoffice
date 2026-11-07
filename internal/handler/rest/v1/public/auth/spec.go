package auth

import "net/http"

type Handler interface {
	// POST /api/login {username, password}
	Login(w http.ResponseWriter, r *http.Request)
}
