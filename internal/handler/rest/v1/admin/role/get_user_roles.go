package role

import (
	"net/http"
	"strings"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"

	roleSvc "github.com/phillong97254/backoffice/internal/service/role"
)

func (i *impl) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimSpace(r.PathValue("userID"))
	if userID == "" {
		apperr.Write(w, apperr.BadRequest("user_id is required"))
		return
	}
	out, err := i.roleSvc.GetUserRoles(r.Context(), roleSvc.UserRoleInput{UserID: userID})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toUserRolesResponse(out))
}
