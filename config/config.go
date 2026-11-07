package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RestServer    RestServer    `mapstructure:"rest_server"`
	API           API           `mapstructure:"api"`
	OAuth         OAuth         `mapstructure:"oauth"`
	WalletService WalletService `mapstructure:"wallet_service"`
	UserService   UserService   `mapstructure:"user_service"`
}

type RestServer struct {
	Name            string `mapstructure:"name"`
	Addr            string `mapstructure:"addr"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"` // giây
}

type API struct {
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type OAuth struct {
	AuthURL      string `mapstructure:"auth_url"`
	Issuer       string `mapstructure:"issuer"`
	Audience     string `mapstructure:"audience"`
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Scopes       string `mapstructure:"scopes"` // cách nhau bằng dấu cách
}

func (o OAuth) JWKSURL() string  { return o.AuthURL + "/.well-known/jwks.json" }
func (o OAuth) TokenURL() string { return o.AuthURL + "/oauth/token" }

type WalletService struct {
	Addr    string `mapstructure:"addr"`
	Timeout int    `mapstructure:"timeout"` // giây
}

type UserService struct {
	URL     string `mapstructure:"url"`
	Timeout int    `mapstructure:"timeout"` // giây
}

// NewConfig đọc config/config.yml, biến môi trường ghi đè theo quy tắc "." -> "_"
// (ví dụ OAUTH_AUTH_URL=http://auth-service:9000, WALLET_SERVICE_ADDR=wallet-service:50051).
func NewConfig() (*Config, error) {
	var cfg Config

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
