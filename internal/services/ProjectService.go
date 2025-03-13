package services

import (
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/websockets"
	"encoding/json"
	"fmt"
	"time"
)

type ProjectService struct {
	ProjectRepo repositories.ProjectRepository
	UserRepo    repositories.UserRepository
	hub         *websockets.Hub
}

func NewProjectService(pr repositories.ProjectRepository, ur repositories.UserRepository, hub *websockets.Hub) *ProjectService {
	return &ProjectService{
		ProjectRepo: pr,
		UserRepo:    ur,
		hub:         hub,
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
	err = ps.ProjectRepo.CreateProject(&project)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}
	project_json, err := json.Marshal(project)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal project: %w", err)
	}
	ps.hub.SendToUser(userId, project_json)

	return &project, nil
}

func (ps *ProjectService) Get(id int) (*models.Project, error) {
	project, err := ps.ProjectRepo.GetProjectById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project by id: %w", err)
	}
	owner, err := ps.UserRepo.GetUserById(project.Owner.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	project.Owner = *owner
	return project, nil
}

func (ps *ProjectService) GetAllProjects(userId int) ([]models.Project, error) {
	projects, err := ps.ProjectRepo.GetAllProjects(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all project: %w", err)
	}
	return projects, nil
}

func (ps *ProjectService) Delete(id int) error {
	err := ps.ProjectRepo.DeleteProject(id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
