package wallet

import (
	"strings"

	"github.com/phil/go-lib/apperr"
)

type AdjustBalanceRequest struct {
	Amount int64  `json:"amount"`
	Note   string `json:"note"`
}

func (r *AdjustBalanceRequest) Validate() error {
	r.Note = strings.TrimSpace(r.Note)
	if r.Amount == 0 {
		return apperr.BadRequest("amount is required")
	}
	return nil
}
