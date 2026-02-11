package main

import (
	"log/slog"
	"os"
)

func main() {
	cnf := config{
		addr: ":5055",
		db:   dbConfig{},
	}

	api := appplication{
		config: cnf,
	}

	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	h := api.mount()
	err := api.run(h)
	if err != nil {
		slog.Error("Error starting the server", "error", err)
		os.Exit(1)
	}
}
