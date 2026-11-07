package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/phil/backoffice/internal/apperr"
	authGwy "github.com/phil/backoffice/internal/gateway/auth"
)

// Login đổi username/password lấy access token ở auth-service.
func (i *impl) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	tok, err := i.gwy.Auth().PasswordToken(ctx, authGwy.PasswordTokenInput{
		Username: input.Username, Password: input.Password,
	})
	if errors.Is(err, authGwy.ErrInvalidCredentials) {
		return LoginOutput{}, apperr.Unauthorized("invalid username or password")
	}
	if err != nil {
		i.logger.Error("login", "error", err)
		return LoginOutput{}, apperr.New(http.StatusBadGateway, "auth-service unavailable")
	}
	return LoginOutput{AccessToken: tok.AccessToken, TokenType: tok.TokenType, ExpiresIn: tok.ExpiresIn, Scope: tok.Scope}, nil
}
