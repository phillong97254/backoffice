package auth

import (
	"strings"

	authSvc "github.com/phil/backoffice/internal/service/auth"
	"github.com/phil/go-lib/apperr"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" || r.Password == "" {
		return apperr.Unauthorized("invalid username or password")
	}
	return nil
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func toLoginResponse(out authSvc.LoginOutput) LoginResponse {
	return LoginResponse{AccessToken: out.AccessToken, TokenType: out.TokenType, ExpiresIn: out.ExpiresIn, Scope: out.Scope}
}
