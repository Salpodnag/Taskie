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
	UserProjectRepo  *repositories.UserProjectRepository
	WebSocketService *websockets.WebSocketService
}

func NewProjectService(pr *repositories.ProjectRepository, ur *repositories.UserRepository, rr *repositories.RoleRepository, upr *repositories.UserProjectRepository, ws *websockets.WebSocketService) *ProjectService {
	return &ProjectService{
		ProjectRepo:      pr,
		UserRepo:         ur,
		RoleRepo:         rr,
		UserProjectRepo:  upr,
		WebSocketService: ws,
	}
}

const (
	Owner       string = "Владелец"
	Participant string = "Участник"
)

func (ps *ProjectService) Create(ProjectDTO dto.CreateProjectDTO) (*dto.ProjectResponseDTO, error) {

	project, err := models.NewProject(ProjectDTO.Name, ProjectDTO.Description, ProjectDTO.Color, ProjectDTO.Privacy)
	if err != nil {
		return nil, err
	}

	err = ps.ProjectRepo.CreateProject(project)
	if err != nil {
		return nil, err
	}

	projectResponseDTO := dto.ProjectToResponseDTO(project)

	if err := ps.WebSocketService.SendMessageBroadcast("project", projectResponseDTO); err != nil {
		return nil, fmt.Errorf("failed to send project message: %w", err)
	}
	err = ps.RoleRepo.CreateDefaultRole(project.Id, Owner)
	if err != nil {
		return nil, err
	}
	err = ps.RoleRepo.CreateDefaultRole(project.Id, Participant)
	if err != nil {
		return nil, err
	}

	// err = ps.UserProjectRepo.AddUserToProject()!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return projectResponseDTO, nil
}

func (ps *ProjectService) GetByIdWOwner(projectID uuid.UUID) (*dto.ProjectResponseDTO, error) {

	project, err := ps.ProjectRepo.GetProjectById(projectID)
	if err != nil {
		return nil, err
	}
	projectResponse := dto.ProjectToResponseDTO(project)
	return projectResponse, nil
}

func (ps *ProjectService) GetAllProjectsWOwner() (*[]dto.ProjectResponseDTO, error) {
	projects, err := ps.ProjectRepo.GetAllProjects()
	if err != nil {
		return nil, fmt.Errorf("failed to get all project: %w", err)
	}
	projectResponseDtos := make([]dto.ProjectResponseDTO, 0, len(projects))
	for _, project := range projects {
		projectResponse := dto.ProjectToResponseDTO(&project)
		projectResponseDtos = append(projectResponseDtos, *projectResponse)
	}

	return &projectResponseDtos, nil
}

func (ps *ProjectService) Delete(ProjectID uuid.UUID) error {

	err := ps.ProjectRepo.DeleteProject(ProjectID)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
