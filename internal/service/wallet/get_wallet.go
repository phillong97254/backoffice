package wallet

import (
	"context"
	"github.com/phillong97254/backoffice/internal/service/grpcerr"

	walletGwy "github.com/phillong97254/backoffice/internal/gateway/wallet"
)

func (i *impl) GetWallet(ctx context.Context, input GetWalletInput) (GetWalletOutput, error) {
	id := resolveAdminID(ctx, input.AdminID)
	ctx, cancel := context.WithTimeout(ctx, i.timeout)
	defer cancel()

	w, err := i.gwy.Wallet().GetWallet(ctx, walletGwy.GetWalletInput{AdminID: id})
	if err != nil {
		return GetWalletOutput{}, grpcerr.FromGRPC(err)
	}
	txs, err := i.gwy.Wallet().ListTransactions(ctx, walletGwy.ListTransactionsInput{AdminID: id, Limit: recentTxLimit})
	if err != nil {
		return GetWalletOutput{}, grpcerr.FromGRPC(err)
	}

	out := GetWalletOutput{Wallet: toWalletOutput(w), Transactions: make([]TransactionOutput, 0, len(txs.Items))}
	for _, t := range txs.Items {
		out.Transactions = append(out.Transactions, TransactionOutput{
			ID: t.ID, Amount: t.Amount, Note: t.Note, CreatedBy: t.CreatedBy, CreatedAtUnix: t.CreatedAtUnix,
		})
	}
	return out, nil
}
