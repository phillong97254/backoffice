package wallet

import (
	"encoding/json"
	"net/http"

	walletSvc "github.com/phillong97254/backoffice/internal/service/wallet"
	"github.com/phillong97254/go-lib/apperr"
	auth "github.com/phillong97254/go-lib/authkit"
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
