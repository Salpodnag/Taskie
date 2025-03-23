package handlers

import (
	"Taskie/internal/services"
	"Taskie/middlewares"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	ProjectService services.ProjectService
}

func NewProjectHandler(ProjectService services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: ProjectService,
	}
}

func (ph *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r)
	if !ok {
		slog.Error("user's id not found in Project creation")
		http.Error(w, "userId not found", http.StatusUnauthorized)
		return
	}

	var reqBody struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		slog.Error("invalid request body", slog.Any("request", r.Body), slog.Any("err", err))
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	if reqBody.Name == "" {
		http.Error(w, fmt.Sprintf("Missing name"), http.StatusBadRequest)
		return
	}
	_, err := ph.ProjectService.Create(reqBody.Name, userID)
	if err != nil {
		slog.Error("failed to create project", slog.String("name", reqBody.Name), slog.Any("err", err))
	}
	w.WriteHeader(http.StatusCreated)
}

func (ph *ProjectHandler) GetById(w http.ResponseWriter, r *http.Request) {

	userID, ok := middlewares.GetUserID(r)
	if !ok {
		slog.Error("user's id not found in request")
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	projectID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("invalid project id format", slog.String("id", projectID.String()), slog.Any("error", err))
		http.Error(w, "Invalid project ID format", http.StatusBadRequest)
		return
	}

	project, err := ph.ProjectService.GetByIdWOwner(projectID, userID)
	if err != nil {
		slog.Error("failed to get project by id: %w", err)
		w.WriteHeader(403)
		return
	}
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		slog.Error("failed to encode project")
		return
	}
}

func (ph *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r)
	if !ok {
		slog.Error("user's id not found in Project creation")
		http.Error(w, "userId not found", http.StatusUnauthorized)
		return
	}
	projects, err := ph.ProjectService.GetAllProjectsWOwner(userID)
	if err != nil {
		slog.Error("failed to get all projects")
		return
	}
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		slog.Error("failed to encode projects")
		return
	}

}

func (ph *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ProjectID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("invalid project id", slog.String("id", ProjectID.String()), slog.Any("error", err))
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	err = ph.ProjectService.Delete(ProjectID)
	if err != nil {
		slog.Error("failed to delete project", slog.String("id", ProjectID.String()), slog.Any("error", err))
		http.Error(w, "Failed to delete project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
}
