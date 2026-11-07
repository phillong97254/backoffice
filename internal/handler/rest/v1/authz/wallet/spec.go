package wallet

import "net/http"

type Handler interface {
	// GET /api/wallets/{adminID}   ("me" = ví của chính mình)
	GetWallet(w http.ResponseWriter, r *http.Request)
	// POST /api/wallets/{adminID}/adjust {amount, note}
	AdjustBalance(w http.ResponseWriter, r *http.Request)
}
