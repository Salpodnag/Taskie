package handlers

import (
	"Taskie/internal/services"
	"Taskie/middlewares"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	userID, ok := middlewares.GetUserID(r)
	var reqBody struct {
		name      string    `json: "name"`
		timeAdded time.Time `json: "timeAdded"`
		CreatedAt time.Time `json: "CreatedAt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

}
