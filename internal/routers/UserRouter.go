package routers

import (
	"Taskie/internal/handlers"
	"Taskie/internal/services"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(userService services.UserService, projectService services.ProjectService) chi.Router {
	r := chi.NewRouter()

	userHandler := handlers.NewUserHandler(userService, projectService)

	r.Get("/", userHandler.AllUsers)
	r.Post("/projects/{projectID}/users/{userID}", userHandler.AddUserToProject)

	return r
}
