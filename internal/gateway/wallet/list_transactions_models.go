package wallet

type ListTransactionsInput struct {
	AdminID string
	Limit   int32
}

type ListTransactionsOutput struct {
	Items []TransactionOutput
}

type TransactionOutput struct {
	ID            string
	Amount        int64
	Note          string
	CreatedBy     string
	CreatedAtUnix int64
}
