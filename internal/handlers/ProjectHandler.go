package handlers

import (
	"Taskie/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectHandler struct {
	ProjectHandler services.ProjectService
}

func NewProjectHandler(ProjectService services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectHandler: ProjectService,
	}
}

func (ph *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		name string `json: "name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// if reqBody.name == "" || reqBody.t == "" || reqBody.Password == "" {
	// 	http.Error(w, "email username and password are required", http.StatusBadRequest)
	// 	return
	// }

}
