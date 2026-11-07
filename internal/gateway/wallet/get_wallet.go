package wallet

import (
	"context"

	walletv1 "github.com/phil/backoffice/pkg/gen/wallet/v1"
)

func (i *impl) GetWallet(ctx context.Context, input GetWalletInput) (WalletOutput, error) {
	w, err := i.client.GetWallet(ctx, &walletv1.GetWalletRequest{AdminId: input.AdminID})
	if err != nil {
		return WalletOutput{}, err
	}
	return toWalletOutput(w), nil
}
