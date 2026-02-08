package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/EagleLizard/ezd-daemon/internal/api"
	"github.com/EagleLizard/ezd-daemon/internal/lib/config"
	"github.com/EagleLizard/ezd-daemon/internal/lib/constants"
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
	os.MkdirAll(constants.LogDir, 0755)
	f, err := os.OpenFile(constants.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	fileLoggerCfg := logging.Config{
		Encoder: zapcore.NewJSONEncoder(logging.GetDefaultEncoderConfig()),
		Writer:  zapcore.AddSync(f),
	}
	logging.Init(loggerCfg, fileLoggerCfg)
	defer logging.Close()

	srv := api.InitServer(cfg, logging.Logger)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}
	api.RunServer(ctx, httpServer)
}
