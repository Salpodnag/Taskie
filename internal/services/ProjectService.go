package services

import (
	"Taskie/internal/dto"
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/websockets"
	"fmt"
	"log/slog"

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

func (ps *ProjectService) Create(userID uuid.UUID, ProjectDTO dto.CreateProjectDTO) (*dto.ProjectResponseDTO, error) {

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
		return nil, err
	}
	slog.Error("или тут ошибка %w", err)

	roleID, err := ps.createDefaultRolesForProject(project.Id)
	if err != nil {
		slog.Error("тут ошибка %w", err)
		return nil, err
	}

	err = ps.UserProjectRepo.AddUserToProject(userID, project.Id, roleID)
	if err != nil {
		slog.Error("нееее тут ошибка %w", err)
		return nil, err
	}
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

func (ps *ProjectService) createDefaultRolesForProject(projectID uuid.UUID) (int, error) {
	roles := []string{Owner, Participant}
	var ownerRoleID int

	for _, role := range roles {
		roleID, err := ps.RoleRepo.GetOrCreateDefaultRole(projectID, role)
		if err != nil {
			return 0, fmt.Errorf("failed to create role '%s': %w", role, err)
		}

		if role == Owner {
			ownerRoleID = roleID
		}
	}

	return ownerRoleID, nil
}
