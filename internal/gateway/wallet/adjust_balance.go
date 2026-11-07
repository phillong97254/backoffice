package wallet

import (
	"context"

	walletv1 "github.com/phillong97254/backoffice/pkg/gen/wallet/v1"
)

func (i *impl) AdjustBalance(ctx context.Context, input AdjustBalanceInput) (WalletOutput, error) {
	w, err := i.client.AdjustBalance(ctx, &walletv1.AdjustBalanceRequest{
		AdminId: input.AdminID, Amount: input.Amount, Note: input.Note,
	})
	if err != nil {
		return WalletOutput{}, err
	}
	return toWalletOutput(w), nil
}
