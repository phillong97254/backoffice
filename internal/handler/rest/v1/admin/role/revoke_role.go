package role

import (
	"net/http"
	"strings"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"

	roleSvc "github.com/phillong97254/backoffice/internal/service/role"
)

func (i *impl) RevokeRole(w http.ResponseWriter, r *http.Request) {
	userID, role := strings.TrimSpace(r.PathValue("userID")), strings.TrimSpace(r.PathValue("role"))
	if userID == "" || role == "" {
		apperr.Write(w, apperr.BadRequest("user_id and role are required"))
		return
	}
	out, err := i.roleSvc.RevokeRole(r.Context(), roleSvc.UserRoleInput{UserID: userID, Role: role})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toUserRolesResponse(out))
}
