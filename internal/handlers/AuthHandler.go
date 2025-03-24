package handlers

import (
	"Taskie/internal/dto"
	"Taskie/internal/models"
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
	if r.Method != http.MethodPost {
		http.Error(w, "need POST method", http.StatusMethodNotAllowed)
		return
	}

	var createUserDto dto.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&createUserDto); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := createUserDto.ValidateCreateUser(); err != nil {
		slog.Error("wrong req body: %w", err)
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
	}

	user, err := ah.authService.Register(createUserDto)
	if err != nil {
		slog.Error("Registration failed", slog.String("email", createUserDto.Email), slog.String("username", createUserDto.Username))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	user, token, err := ah.authService.Login(reqBody.Identifier, reqBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	response := struct {
		User *models.User `json:"user"`
		// Token string       `json:"token"`
	}{
		User: user,
		// Token: token,
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "set-token",
		Value: token,
		Path:  "/",
	})
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
