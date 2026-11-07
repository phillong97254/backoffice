package wallet

import walletv1 "github.com/phil/backoffice/pkg/gen/wallet/v1"

type WalletOutput struct {
	AdminID       string
	Currency      string
	Balance       int64
	UpdatedAtUnix int64
}

func toWalletOutput(w *walletv1.Wallet) WalletOutput {
	return WalletOutput{
		AdminID:       w.GetAdminId(),
		Currency:      w.GetCurrency(),
		Balance:       w.GetBalance(),
		UpdatedAtUnix: w.GetUpdatedAtUnix(),
	}
}
