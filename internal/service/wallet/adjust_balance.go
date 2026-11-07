package wallet

import (
	"context"

	walletGwy "github.com/phil/backoffice/internal/gateway/wallet"
	"github.com/phil/go-lib/apperr"
)

func (i *impl) AdjustBalance(ctx context.Context, input AdjustBalanceInput) (WalletOutput, error) {
	if input.Amount == 0 {
		return WalletOutput{}, apperr.BadRequest("amount is required")
	}
	ctx, cancel := context.WithTimeout(ctx, i.timeout)
	defer cancel()

	w, err := i.gwy.Wallet().AdjustBalance(ctx, walletGwy.AdjustBalanceInput{
		AdminID: resolveAdminID(ctx, input.AdminID), Amount: input.Amount, Note: input.Note,
	})
	if err != nil {
		return WalletOutput{}, fromGRPC(err)
	}
	return toWalletOutput(w), nil
}
