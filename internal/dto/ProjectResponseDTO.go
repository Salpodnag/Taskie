package dto

import (
	"Taskie/internal/models"
	"time"

	"github.com/google/uuid"
)

type ProjectResponseDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Privacy     string    `json:"privacy"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ProjectToResponseDTO(project *models.Project) *ProjectResponseDTO {
	return &ProjectResponseDTO{
		Id:          project.Id,
		Name:        project.Name,
		Description: project.Description,
		Color:       project.Color,
		Privacy:     project.Privacy,
		CreatedAt:   project.CreatedAt,
	}
}
