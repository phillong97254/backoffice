package auth

import (
	"net/http"
	"time"

	"github.com/phil/backoffice/config"
)

type impl struct {
	tokenURL string
	http     *http.Client
}

func New(cfg *config.Config) Gateway {
	return &impl{
		tokenURL: cfg.OAuth.TokenURL(),
		http:     &http.Client{Timeout: 5 * time.Second},
	}
}
