package services

import (
	"Taskie/internal/repositories"

	"github.com/google/uuid"
)

type RoleService struct {
	ProjectRepo     *repositories.ProjectRepository
	UserRepo        *repositories.UserRepository
	RoleRepo        *repositories.RoleRepository
	UserProjectRepo *repositories.UserProjectRepository
}

func NewRoleService(pr *repositories.ProjectRepository, ur *repositories.UserRepository, rr *repositories.RoleRepository, upr *repositories.UserProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepo:     pr,
		UserRepo:        ur,
		RoleRepo:        rr,
		UserProjectRepo: upr,
	}
}

const (
	Owner       string = "Владелец"
	Participant string = "Участник"
)

func (rs *ProjectService) createDefaultRolesForProject(projectID uuid.UUID) (int, error) {

	roleID, err := rs.RoleRepo.GetOrCreateDefaultRole(projectID, Owner)
	if err != nil {
		return 0, err
	}

	ownerRoleID := roleID

	return ownerRoleID, nil
}

func (rs *ProjectService) CreateRoleParticipant(projectID uuid.UUID) (int, error) {
	roleID, err := rs.RoleRepo.CreateRole(projectID, Participant)
	if err != nil {
		return 0, err
	}
	return roleID, nil
}
