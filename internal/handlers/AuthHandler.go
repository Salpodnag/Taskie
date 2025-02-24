package handlers

import (
	"Taskie/internal/services"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	err := ah.authService.Register(reqBody.Email, reqBody.Username, reqBody.Password)
	if err != nil {
		slog.Error("Registration failed", slog.String("email", reqBody.Email), slog.String("username", reqBody.Username))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))

}
