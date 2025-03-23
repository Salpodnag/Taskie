package services

import (
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/websockets"
	"fmt"
	"time"
)

type ProjectService struct {
	ProjectRepo      *repositories.ProjectRepository
	UserRepo         *repositories.UserRepository
	RoleRepo         *repositories.RoleRepository
	WebSocketService *websockets.WebSocketService
}

func NewProjectService(pr *repositories.ProjectRepository, ur *repositories.UserRepository, rr *repositories.RoleRepository, ws *websockets.WebSocketService) *ProjectService {
	return &ProjectService{
		ProjectRepo:      pr,
		UserRepo:         ur,
		RoleRepo:         rr,
		WebSocketService: ws,
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

	if err := ps.WebSocketService.SendMessageBroadcast("project", project); err != nil {
		return nil, fmt.Errorf("failed to send project message: %w", err)
	}
	err = ps.RoleRepo.CreateDefaultRoles(project.Id)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (ps *ProjectService) GetByIdWOwner(id int, userID int) (*models.Project, error) {

	project, err := ps.ProjectRepo.GetProjectById(id)
	if err != nil {
		return nil, err
	}
	if project.Owner.Id != userID {
		return nil, fmt.Errorf("nuh-uh, не твой проект: %d", userID)
	}
	return project, nil
}

func (ps *ProjectService) GetAllProjectsWOwner(userId int) ([]models.Project, error) {
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
