package wallet

import walletSvc "github.com/phillong97254/backoffice/internal/service/wallet"

type WalletResponse struct {
	AdminID       string `json:"admin_id"`
	Currency      string `json:"currency"`
	Balance       int64  `json:"balance"`
	UpdatedAtUnix int64  `json:"updated_at_unix"`
}

func toWalletResponse(w walletSvc.WalletOutput) WalletResponse {
	return WalletResponse{AdminID: w.AdminID, Currency: w.Currency, Balance: w.Balance, UpdatedAtUnix: w.UpdatedAtUnix}
}
