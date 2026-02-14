package main

import (
	"api-ecom/internal/env"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	cnf := config{
		addr: ":5055",
		db: dbConfig{
			dsn: env.Get("GOOSE_DBSTRING", ""),
		},
	}

	conn, err := pgx.Connect(ctx, env.Get("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=api-ecom-go sslmode=disable"))
	if err != nil {
		slog.Error("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	slog.Info("connected to database", "dsn", cnf.db.dsn)

	api := appplication{
		config: cnf,
		db:     conn,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Error starting the server", "error", err)
		os.Exit(1)
	}
}
