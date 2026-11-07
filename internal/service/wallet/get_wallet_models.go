package wallet

type GetWalletInput struct {
	AdminID string
}

type GetWalletOutput struct {
	Wallet       WalletOutput
	Transactions []TransactionOutput
}

type TransactionOutput struct {
	ID            string
	Amount        int64
	Note          string
	CreatedBy     string
	CreatedAtUnix int64
}
