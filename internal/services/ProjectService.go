package services

import (
	"Taskie/internal/dto"
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/websockets"
	"fmt"

	"github.com/google/uuid"
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

func (ps *ProjectService) Create(ProjectDTO dto.CreateProjectDTO, userID uuid.UUID) (*dto.ProjectResponseDTO, error) {
	owner, err := ps.UserRepo.GetUserById(userID)
	if err != nil {
		return nil, err
	}
	project, err := models.NewProject(ProjectDTO.Name, *owner, ProjectDTO.Description, ProjectDTO.Color, ProjectDTO.Privacy)
	if err != nil {
		return nil, err
	}

	err = ps.ProjectRepo.CreateProject(project)
	if err != nil {
		return nil, err
	}

	projectResponseDTO := dto.ProjectToResponseDTO(project, owner)

	if err := ps.WebSocketService.SendMessageBroadcast("project", projectResponseDTO); err != nil {
		return nil, fmt.Errorf("failed to send project message: %w", err)
	}
	err = ps.RoleRepo.CreateDefaultRoles(project.Id)
	if err != nil {
		return nil, err
	}

	return projectResponseDTO, nil
}

func (ps *ProjectService) GetByIdWOwner(projectID uuid.UUID, userID uuid.UUID) (*models.Project, error) {

	project, err := ps.ProjectRepo.GetProjectById(projectID)
	if err != nil {
		return nil, err
	}
	if project.Owner.Id != userID {
		return nil, fmt.Errorf("nuh-uh, не твой проект: %d", userID)
	}
	return project, nil
}

func (ps *ProjectService) GetAllProjectsWOwner(userID uuid.UUID) ([]models.Project, error) {
	projects, err := ps.ProjectRepo.GetAllProjects(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all project: %w", err)
	}
	return projects, nil
}

func (ps *ProjectService) Delete(ProjectID uuid.UUID) error {

	err := ps.ProjectRepo.DeleteProject(ProjectID)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
