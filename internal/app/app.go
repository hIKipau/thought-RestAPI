package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"thought-RestAPI/internal/adapter/postgresql"
	"thought-RestAPI/internal/config"
	httpRouter "thought-RestAPI/internal/transport/http"
	"thought-RestAPI/internal/usecase"
	"time"
)

func Run(ctx context.Context, cfg *config.Config, log *slog.Logger) error {
	const op = "internal/app/Run"

	pgsql, err := postgresql.New(ctx, cfg.DatabaseURL, log)
	if err != nil {
		return fmt.Errorf("%s: Failed to connect to database. Error: %s", op, err.Error())
	}

	thought := usecase.NewThought(pgsql)

	router := httpRouter.Router(thought)
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // ждём здесь сигнала (Ctrl+C или SIGTERM)

	log.Info("Interrupt received, shutting down...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctxShutdown)
	pgsql.Close()

	return nil
}
