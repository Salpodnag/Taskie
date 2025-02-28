package services

import (
	"Taskie/internal/repositories"
)

type ProjectService struct {
	ProjectRepo repositories.ProjectRepository
}

func NewProjectService(pr repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepo: pr,
	}
}

// func (pr *ProjectService) ProjectCreation() error {
// 	var project models.Project

// }
