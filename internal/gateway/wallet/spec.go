package wallet

import "context"

// Gateway gọi wallet-service qua gRPC. Lỗi trả về là gRPC status gốc (không wrap)
// để tầng service đọc được status code.
type Gateway interface {
	GetWallet(ctx context.Context, input GetWalletInput) (WalletOutput, error)
	ListTransactions(ctx context.Context, input ListTransactionsInput) (ListTransactionsOutput, error)
	AdjustBalance(ctx context.Context, input AdjustBalanceInput) (WalletOutput, error)
	Close() error
}
