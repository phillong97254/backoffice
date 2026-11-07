package wallet

import walletSvc "github.com/phil/backoffice/internal/service/wallet"

type GetWalletResponse struct {
	Wallet       WalletResponse        `json:"wallet"`
	Transactions []TransactionResponse `json:"transactions"`
}

type TransactionResponse struct {
	ID            string `json:"id"`
	Amount        int64  `json:"amount"`
	Note          string `json:"note"`
	CreatedBy     string `json:"created_by"`
	CreatedAtUnix int64  `json:"created_at_unix"`
}

func toGetWalletResponse(out walletSvc.GetWalletOutput) GetWalletResponse {
	resp := GetWalletResponse{Wallet: toWalletResponse(out.Wallet), Transactions: make([]TransactionResponse, 0, len(out.Transactions))}
	for _, t := range out.Transactions {
		resp.Transactions = append(resp.Transactions, TransactionResponse{
			ID: t.ID, Amount: t.Amount, Note: t.Note, CreatedBy: t.CreatedBy, CreatedAtUnix: t.CreatedAtUnix,
		})
	}
	return resp
}
