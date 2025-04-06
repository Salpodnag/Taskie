package handlers

import (
	"Taskie/internal/services"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService    services.UserService
	ProjectService services.ProjectService
}

func NewUserHandler(UserService services.UserService, ProjectService services.ProjectService) *UserHandler {
	return &UserHandler{
		UserService:    UserService,
		ProjectService: ProjectService,
	}
}

func (uh *UserHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.UserService.AllUsers()
	if err != nil {
		slog.Error("failed to get all users", "error", err)
		http.Error(w, "failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("failed to encode users response", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) AddUserToProject(w http.ResponseWriter, r *http.Request) {

	projectIDStr := chi.URLParam(r, "projectID")
	userIDStr := chi.URLParam(r, "userID")

	projectID, err := uuid.Parse(projectIDStr)
	if err != nil {
		http.Error(w, "invalid project ID", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.ProjectService.AddUserToProject(userID, projectID)
	if err != nil {
		slog.Error("failed to add user to project", "error", err)
		http.Error(w, "failed to add user to project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
