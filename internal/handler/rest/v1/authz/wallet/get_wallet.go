package wallet

import (
	"net/http"

	walletSvc "github.com/phil/backoffice/internal/service/wallet"
	"github.com/phil/go-lib/apperr"
	auth "github.com/phil/go-lib/authkit"
)

func (i *impl) GetWallet(w http.ResponseWriter, r *http.Request) {
	out, err := i.walletSvc.GetWallet(r.Context(), walletSvc.GetWalletInput{AdminID: r.PathValue("adminID")})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toGetWalletResponse(out))
}
