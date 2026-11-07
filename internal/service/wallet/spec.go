package wallet

import "context"

type Service interface {
	// GetWallet: thông tin ví + 20 giao dịch gần nhất. AdminID "me" = ví của chính mình.
	GetWallet(ctx context.Context, input GetWalletInput) (GetWalletOutput, error)
	// AdjustBalance: cộng/trừ số dư (wallet-service chỉ cho superadmin).
	AdjustBalance(ctx context.Context, input AdjustBalanceInput) (WalletOutput, error)
}
