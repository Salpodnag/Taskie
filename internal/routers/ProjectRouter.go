package routers

import (
	"Taskie/internal/handlers"
	"Taskie/internal/services"

	"github.com/go-chi/chi/v5"
)

func NewProjectRouter(projectService services.ProjectService) chi.Router {
	r := chi.NewRouter()

	projectHandler := handlers.NewProjectHandler(projectService)

	r.Post("/", projectHandler.Create)

	return r
}
