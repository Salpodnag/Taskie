package routers

import (
	"Taskie/internal/handlers"
	"Taskie/internal/services"

	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(authService services.AuthService) chi.Router {
	r := chi.NewRouter()

	authHandler := handlers.NewAuthHandler(authService)

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	return r
}
