package router

import (
	"Taskie/internal/handlers"
	"Taskie/internal/services"
	"net/http"
)

func NewRouter(authService services.AuthService) http.Handler {
	authHandler := handlers.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)
	return mux
}
