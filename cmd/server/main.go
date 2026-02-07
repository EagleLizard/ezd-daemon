package main

import (
	"context"
	"net"
	"net/http"

	"github.com/EagleLizard/ezd-daemon/internal/api"
	"github.com/EagleLizard/ezd-daemon/internal/lib/config"
	"github.com/EagleLizard/ezd-daemon/internal/lib/logging"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()
	cfg := config.EzdDConfig
	startServer(ctx, cfg)
}

func startServer(ctx context.Context, cfg *config.EzdDConfigType) {
	loggerCfg := logging.Config{
		Encoder: zapcore.NewJSONEncoder(logging.GetDefaultEncoderConfig()),
	}
	logging.Init(loggerCfg)
	defer logging.Close()

	srv := api.InitServer(cfg, logging.Logger)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}
	api.RunServer(ctx, httpServer)
}
