package wallet

import (
	"context"

	auth "github.com/phil/backoffice/pkg/authkit"
)

// resolveAdminID: "me" = admin đang đăng nhập.
func resolveAdminID(ctx context.Context, id string) string {
	if id == meAlias {
		if c, ok := auth.ClaimsFrom(ctx); ok {
			return c.Subject
		}
	}
	return id
}
