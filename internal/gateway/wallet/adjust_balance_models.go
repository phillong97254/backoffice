package wallet

type AdjustBalanceInput struct {
	AdminID string
	Amount  int64
	Note    string
}
