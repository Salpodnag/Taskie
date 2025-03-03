package main

import (
	"Taskie/cfg"
	"Taskie/db"
	"Taskie/internal/repositories"
	"Taskie/internal/routers"
	"Taskie/internal/services"
	"Taskie/logger"
	"Taskie/middlewares"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
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
	projectRepository := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(*projectRepository)

	r := chi.NewRouter()

	r.Use(middlewares.CorsMiddleware)

	r.Mount("/auth", routers.NewAuthRouter(*authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware)
		r.Mount("/project", routers.NewProjectRouter(*projectService))
	})

	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
