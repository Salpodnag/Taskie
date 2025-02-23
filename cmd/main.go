package main

import (
	"Taskie/cfg"
	"Taskie/db"
	"Taskie/internal/repositories"
	"Taskie/internal/services"
	"Taskie/logger"
	"Taskie/router"
	"fmt"
	"log"
	"log/slog"
	"net/http"
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
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(cfg.JWT, *userRepository)
	r := router.NewRouter(*authService)
	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
