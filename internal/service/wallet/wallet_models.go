package wallet

import walletGwy "github.com/phillong97254/backoffice/internal/gateway/wallet"

type WalletOutput struct {
	AdminID       string
	Currency      string
	Balance       int64
	UpdatedAtUnix int64
}

func toWalletOutput(w walletGwy.WalletOutput) WalletOutput {
	return WalletOutput{AdminID: w.AdminID, Currency: w.Currency, Balance: w.Balance, UpdatedAtUnix: w.UpdatedAtUnix}
}
