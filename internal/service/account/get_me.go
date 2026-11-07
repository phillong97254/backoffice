package account

import (
	"context"
	"net/http"

	"github.com/phil/backoffice/internal/apperr"
	userGwy "github.com/phil/backoffice/internal/gateway/user"
	auth "github.com/phil/backoffice/pkg/authkit"
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
