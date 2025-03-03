package handlers

import (
	"Taskie/internal/services"
	"Taskie/middlewares"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
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
	}
	project, err := ph.ProjectService.Create(reqBody.Name, userID)
	if err != nil {
		slog.Error("failed to create project", slog.String("name", reqBody.Name), slog.Any("err", err))
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (ph *ProjectHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var id int
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		slog.Error("failed to get project ID from response", slog.String("id", string(id)))
		return
	}
	if id == 0 {
		slog.Error("missing id", slog.String("id", string(id)))
	}
	project, err := ph.ProjectService.Get(id)
	if err != nil {
		slog.Error("failed to get project by id", slog.String("id", string(id)))
		return
	}
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		slog.Error("failed to encode project")
		return
	}
}
