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
	r.Get("/{id}", projectHandler.GetById)
	r.Get("/", projectHandler.GetAllProjects)
	r.Delete("/{id}", projectHandler.Delete)

	return r
}
