package wallet

import (
	"context"

	auth "github.com/phillong97254/go-lib/authkit"
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
