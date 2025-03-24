package dto

import (
	"Taskie/internal/models"

	"github.com/google/uuid"
)

type ProjectResponseDTO struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Color       string      `json:"color"`
	Privacy     string      `json:"privacy"`
	Owner       models.User `json:"user"`
}

func ProjectToResponseDTO(project *models.Project, owner *models.User) *ProjectResponseDTO {
	return &ProjectResponseDTO{
		Id:          project.Id,
		Name:        project.Name,
		Description: project.Description,
		Color:       project.Color,
		Privacy:     project.Privacy,
		Owner:       *owner,
	}
}
