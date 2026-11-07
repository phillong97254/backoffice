package role

import (
	"net/http"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"
)

func (i *impl) ListRoles(w http.ResponseWriter, r *http.Request) {
	out, err := i.roleSvc.ListRoles(r.Context())
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toListRolesResponse(out))
}
