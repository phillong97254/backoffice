// backoffice: REST API (BFF) + trang admin backoffice.
//
//	Browser --REST+JWT--> backoffice
//	backoffice --gRPC (forward JWT của admin)--> wallet-service
//	backoffice --REST (JWT của chính backoffice, client_credentials)--> user-service
package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phillong97254/backoffice/cmd/rest/routes"
	"github.com/phillong97254/backoffice/config"
	"github.com/phillong97254/backoffice/internal/gateway"
	v1 "github.com/phillong97254/backoffice/internal/handler/rest/v1"
	"github.com/phillong97254/backoffice/web"
	auth "github.com/phillong97254/go-lib/authkit"
)

func main() {
	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("server", cfg.RestServer.Name, "env", cfg.API.Env)
	slog.SetDefault(logger)

	// Initialize Gateway (client gọi sang auth-service, user-service, wallet-service)
	gwy, err := gateway.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer gwy.Close()

	// Authorize: verify JWT bằng JWKS của auth-service
	verifier := auth.NewVerifier(cfg.OAuth.JWKSURL(), cfg.OAuth.Issuer, cfg.OAuth.Audience)

	// Handler (service được tạo bên trong theo từng nhóm handler)
	handler := v1.New(cfg, logger, gwy)

	// Register Routes
	mux := http.NewServeMux()
	routes.V1Routes(mux, verifier, handler)
	routes.WebRoutes(mux, web.FS)

	// Start Server
	srv := &http.Server{Addr: cfg.RestServer.Addr, Handler: routes.LogRequests(logger, mux), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		logger.Info("backoffice listening", "addr", cfg.RestServer.Addr, "open", "http://localhost"+cfg.RestServer.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.RestServer.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "error", err)
	}
	logger.Info("backoffice stopped")
}
