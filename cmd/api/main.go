package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"thoughtRestApi/internal/config"
	"thoughtRestApi/internal/logger"
	"thoughtRestApi/internal/repository/postgresql"
)

func main() {
	//TODO: init router
	//TODO: start app
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logs := logger.New(cfg.Env)

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		logs.Error("Error connecting to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	postgreSQL := postgresql.New(conn)

}
