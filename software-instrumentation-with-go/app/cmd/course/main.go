package main

import (
	"app/course/server/apiserver"
	"app/internal/config"
	"app/internal/database"
	"app/internal/logger"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "./configs/config.yaml", "Configuration file path.")
	flag.Parse()

	logger.Init()
	defer logger.Log.Sync()

	log := logger.Log

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal("failed to load config",
			zap.String("config", *configPath),
			zap.Error(err),
		)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("failed to connect to database",
			zap.Error(err),
		)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	server := apiserver.New(&apiserver.Option{
		Cfg: cfg,
		DB:  db,
		Log: log,
	})

	log.Info("server starting")

	if err := server.Run(ctx); err != nil {
		log.Fatal("server stopped with error",
			zap.Error(err),
		)
	}

	log.Info("server stopped gracefully")
}
