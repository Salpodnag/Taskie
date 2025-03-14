package handlers

import (
	"Taskie/internal/services"
	"Taskie/middlewares"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("invalid id format", slog.String("id", strconv.Itoa(id)))
		return
	}
	if id == 0 {
		slog.Error("missing id", slog.String("id", strconv.Itoa(id)))
	}
	project, err := ph.ProjectService.Get(id)
	if err != nil {
		slog.Error("failed to get project by id", slog.String("id", strconv.Itoa(id)))
		return
	}
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		slog.Error("failed to encode project")
		return
	}
}

func (ph *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r)
	if !ok {
		slog.Error("user's id not found in Project creation")
		http.Error(w, "userId not found", http.StatusUnauthorized)
		return
	}
	_, err := ph.ProjectService.GetAllProjects(userId)
	if err != nil {
		slog.Error("failed to get all projects")
		return
	}
	w.WriteHeader(200)

}

func (ph *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("invalid id format", slog.String("id", strconv.Itoa(id)))
		return
	}
	if id == 0 {
		slog.Error("missing id", slog.String("id", strconv.Itoa(id)))
	}
	err = ph.ProjectService.Delete(id)
	if err != nil {
		slog.Error("failed to delete project", slog.String("id", strconv.Itoa(id)))
		return
	}
	w.WriteHeader(204)
}
