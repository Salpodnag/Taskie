package main

import (
	"Taskie/cfg"
	"Taskie/db"
	"Taskie/logger"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	lg := logger.New()
	slog.SetDefault(lg)

	cfg, err := cfg.Load()
	if err != nil {
		lg.Error(
			"failed to load config",
			slog.String("error", err.Error()),
		)
	}

	db, err := db.NewClient(cfg)
	if err != nil {
		lg.Error(
			"failed to connect to database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
	fmt.Println(db)
}
