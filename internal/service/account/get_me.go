package account

import (
	"context"
	"net/http"

	userGwy "github.com/phil/backoffice/internal/gateway/user"
	"github.com/phil/go-lib/apperr"
	auth "github.com/phil/go-lib/authkit"
)

func (i *impl) GetMe(ctx context.Context) (GetMeOutput, error) {
	c, ok := auth.ClaimsFrom(ctx)
	if !ok {
		return GetMeOutput{}, apperr.Unauthorized("unauthenticated")
	}
	p, err := i.gwy.User().GetProfile(ctx, userGwy.GetProfileInput{ID: c.Subject})
	if err != nil {
		i.logger.Error("get profile", "sub", c.Subject, "error", err)
		return GetMeOutput{}, apperr.New(http.StatusBadGateway, "user-service: "+err.Error())
	}
	return GetMeOutput{
		Subject: c.Subject,
		Roles:   c.Roles,
		Scope:   c.Scope,
		Profile: ProfileOutput{ID: p.ID, Name: p.Name, Email: p.Email, Department: p.Department},
	}, nil
}
