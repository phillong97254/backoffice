package auth

import (
	"encoding/json"
	"net/http"

	"github.com/phil/backoffice/internal/apperr"
	authSvc "github.com/phil/backoffice/internal/service/auth"
	authkit "github.com/phil/backoffice/pkg/authkit"
)

func (i *impl) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apperr.Write(w, apperr.BadRequest("bad json"))
		return
	}
	if err := req.Validate(); err != nil {
		apperr.Write(w, err)
		return
	}

	out, err := i.authSvc.Login(r.Context(), authSvc.LoginInput{Username: req.Username, Password: req.Password})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	authkit.WriteJSON(w, http.StatusOK, toLoginResponse(out))
}
