package authz

import (
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz/account"
	"github.com/phillong97254/backoffice/internal/handler/rest/v1/authz/wallet"
)

// Handler cho các endpoint cần access token của người dùng.
type Handler interface {
	GetAccountHandler() account.Handler
	GetWalletHandler() wallet.Handler
}
