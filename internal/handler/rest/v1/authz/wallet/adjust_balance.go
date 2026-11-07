package wallet

import (
	"encoding/json"
	"net/http"

	"github.com/phil/backoffice/internal/apperr"
	walletSvc "github.com/phil/backoffice/internal/service/wallet"
	auth "github.com/phil/backoffice/pkg/authkit"
)

func (i *impl) AdjustBalance(w http.ResponseWriter, r *http.Request) {
	var req AdjustBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apperr.Write(w, apperr.BadRequest("amount is required"))
		return
	}
	if err := req.Validate(); err != nil {
		apperr.Write(w, err)
		return
	}

	out, err := i.walletSvc.AdjustBalance(r.Context(), walletSvc.AdjustBalanceInput{
		AdminID: r.PathValue("adminID"), Amount: req.Amount, Note: req.Note,
	})
	if err != nil {
		apperr.Write(w, err)
		return
	}
	auth.WriteJSON(w, http.StatusOK, toWalletResponse(out))
}
