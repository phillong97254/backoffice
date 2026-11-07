package account

import (
	"net/http"

	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"
)

func (i *impl) GetMe(w http.ResponseWriter, r *http.Request) {
	out, err := i.accountSvc.GetMe(r.Context())
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toGetMeResponse(out))
}
