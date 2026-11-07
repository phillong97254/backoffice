package user

import (
	"context"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/phil/backoffice/config"
)

type impl struct {
	baseURL string
	http    *http.Client // tự xin + tự refresh token client_credentials của gateway
}

func New(cfg *config.Config) Gateway {
	cc := clientcredentials.Config{
		ClientID:     cfg.OAuth.ClientID,
		ClientSecret: cfg.OAuth.ClientSecret,
		TokenURL:     cfg.OAuth.TokenURL(),
		Scopes:       strings.Fields(cfg.OAuth.Scopes),
	}
	client := cc.Client(context.Background())
	client.Timeout = time.Duration(cfg.UserService.Timeout) * time.Second
	return &impl{baseURL: cfg.UserService.URL, http: client}
}
