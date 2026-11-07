package role

import (
	"strings"

	"github.com/phillong97254/go-lib/apperr"
)

type GrantRoleRequest struct {
	UserID string `json:"-"` // từ path
	Role   string `json:"role"`
}

func (r *GrantRoleRequest) Validate() error {
	r.UserID = strings.TrimSpace(r.UserID)
	r.Role = strings.TrimSpace(r.Role)
	if r.UserID == "" || r.Role == "" {
		return apperr.BadRequest("user_id and role are required")
	}
	return nil
}
