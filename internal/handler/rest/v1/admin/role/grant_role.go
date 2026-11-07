package role

import (
	"encoding/json"
	"net/http"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"

	roleSvc "github.com/phillong97254/backoffice/internal/service/role"
)

func (i *impl) GrantRole(w http.ResponseWriter, r *http.Request) {
	var req GrantRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apperr.Write(w, apperr.BadRequest("bad json"))
		return
	}
	req.UserID = r.PathValue("userID")
	if err := req.Validate(); err != nil {
		apperr.Write(w, err)
		return
	}
	out, err := i.roleSvc.GrantRole(r.Context(), roleSvc.UserRoleInput{UserID: req.UserID, Role: req.Role})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toUserRolesResponse(out))
}
