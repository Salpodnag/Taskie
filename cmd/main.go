package main

import (
	"Taskie/cfg"
	"Taskie/db"
	"Taskie/internal/repositories"
	"Taskie/internal/routers"
	"Taskie/internal/services"
	"Taskie/logger"
	"Taskie/middlewares"
	"Taskie/websockets"
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

	hub := websockets.NewHub()
	wsService := websockets.NewWebSocketService(hub)

	userRepository := repositories.NewUserRepository(db)

	roleRepository := repositories.NewRoleRepository(db)
	userProjectRepository := repositories.NewUserProjectRepository(db)

	projectRepository := repositories.NewProjectRepository(db)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(cfg.JWT, userRepository, projectRepository, wsService)
	projectService := services.NewProjectService(projectRepository, userRepository, roleRepository, userProjectRepository, wsService)

	r := chi.NewRouter()

	r.Use(middlewares.CorsMiddleware)
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		websockets.WsHandler(w, r, hub)
	})
	r.Mount("/auth", routers.NewAuthRouter(*authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware)
		r.Mount("/projects", routers.NewProjectRouter(*projectService))
		r.Mount("/users", routers.NewUserRouter(*userService, *projectService))
	})

	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
