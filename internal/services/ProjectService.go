package services

import (
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"fmt"
	"time"
)

type ProjectService struct {
	ProjectRepo repositories.ProjectRepository
	UserRepo    repositories.UserRepository
}

func NewProjectService(pr repositories.ProjectRepository, ur repositories.UserRepository) *ProjectService {
	return &ProjectService{
		ProjectRepo: pr,
		UserRepo:    ur,
	}
}

func (ps *ProjectService) Create(name string, userId int) (*models.Project, error) {
	var project models.Project
	project.Name = name
	project.CreatedAt = time.Now()
	owner, err := ps.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id %d: %w", userId, err)
	}
	project.Owner = *owner
	return &project, nil
}
