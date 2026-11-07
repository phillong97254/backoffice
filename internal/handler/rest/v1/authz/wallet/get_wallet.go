package wallet

import (
	"net/http"

	"github.com/phil/backoffice/internal/apperr"
	walletSvc "github.com/phil/backoffice/internal/service/wallet"
	auth "github.com/phil/backoffice/pkg/authkit"
)

func (i *impl) GetWallet(w http.ResponseWriter, r *http.Request) {
	out, err := i.walletSvc.GetWallet(r.Context(), walletSvc.GetWalletInput{AdminID: r.PathValue("adminID")})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toGetWalletResponse(out))
}
