package public

import "github.com/phil/backoffice/internal/handler/rest/v1/public/auth"

// Handler cho các endpoint không cần access token.
type Handler interface {
	GetAuthHandler() auth.Handler
}
