package services

import (
	"Taskie/internal/dto"
	"Taskie/internal/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	ProjectRepo     *repositories.ProjectRepository
	UserRepo        *repositories.UserRepository
	RoleRepo        *repositories.RoleRepository
	UserProjectRepo *repositories.UserProjectRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (us *UserService) AllUsers() ([]dto.UserResponseDTO, error) {
	users, err := us.UserRepo.AllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]dto.UserResponseDTO, 0, len(users))
	for _, user := range users {
		userDTO := dto.UserToResponseDTO(user)
		userDTOs = append(userDTOs, *userDTO)
	}

	return userDTOs, nil
}

func (ps *ProjectService) AddUserToProject(userID uuid.UUID, projectID uuid.UUID) error {
	roleID, err := ps.CreateRoleParticipant(projectID)
	if err != nil {
		return err
	}
	err = ps.UserProjectRepo.AddUserToProject(userID, projectID, roleID)
	if err != nil {
		return err
	}
	return nil
}
