package wallet

import (
	"net/http"

	walletSvc "github.com/phillong97254/backoffice/internal/service/wallet"
	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"
)

func (i *impl) GetWallet(w http.ResponseWriter, r *http.Request) {
	out, err := i.walletSvc.GetWallet(r.Context(), walletSvc.GetWalletInput{AdminID: r.PathValue("adminID")})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toGetWalletResponse(out))
}
