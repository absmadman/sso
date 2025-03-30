package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
)

func main() {
	// init config
	cfg := config.MustRead()

	// TODO: init logger
	logger := initLogger(cfg.Env)

	logger.Info("init logger")

	// app.New(logger);
	// TODO: init app

	// TODO: gRPC serve
}

func initLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "local":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "production":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
