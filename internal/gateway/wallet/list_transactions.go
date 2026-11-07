package wallet

import (
	"context"

	walletv1 "github.com/phil/backoffice/pkg/gen/wallet/v1"
)

func (i *impl) ListTransactions(ctx context.Context, input ListTransactionsInput) (ListTransactionsOutput, error) {
	resp, err := i.client.ListTransactions(ctx, &walletv1.ListTransactionsRequest{AdminId: input.AdminID, Limit: input.Limit})
	if err != nil {
		return ListTransactionsOutput{}, err
	}
	out := ListTransactionsOutput{Items: make([]TransactionOutput, 0, len(resp.GetItems()))}
	for _, t := range resp.GetItems() {
		out.Items = append(out.Items, TransactionOutput{
			ID: t.GetId(), Amount: t.GetAmount(), Note: t.GetNote(),
			CreatedBy: t.GetCreatedBy(), CreatedAtUnix: t.GetCreatedAtUnix(),
		})
	}
	return out, nil
}
