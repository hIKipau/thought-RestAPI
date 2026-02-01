package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"thought-RestAPI/internal/app"
	"thought-RestAPI/internal/config"
	"thought-RestAPI/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.New(cfg.Env)

	err = app.Run(context.Background(), cfg, log)
	if err != nil {
		log.Error("Could not start application", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
