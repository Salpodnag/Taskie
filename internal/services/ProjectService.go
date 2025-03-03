package services

import (
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/middlewares"
	"time"
)

type ProjectService struct {
	ProjectRepo repositories.ProjectRepository
}

func NewProjectService(pr repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepo: pr,
	}
}

func (ps *ProjectService) Create(name string, userId int) error {
	var project models.Project
	project.Name = name
	project.CreatedAt = time.Now()
	project.Owner = repositories.UserRepository.GetUserById(userId)

}

// }
