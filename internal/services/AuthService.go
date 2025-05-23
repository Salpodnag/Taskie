package services

import (
	"Taskie/cfg"
	"Taskie/internal/dto"
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/internal/utils"
	"Taskie/websockets"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	JwtKey           cfg.JWT
	UserRepo         *repositories.UserRepository
	ProjectRepo      *repositories.ProjectRepository
	WebSocketService *websockets.WebSocketService
}

func NewAuthService(JWT cfg.JWT, ur *repositories.UserRepository, ProjectRepo *repositories.ProjectRepository, WebSocketService *websockets.WebSocketService) *AuthService {
	return &AuthService{
		JwtKey:           JWT,
		UserRepo:         ur,
		ProjectRepo:      ProjectRepo,
		WebSocketService: WebSocketService,
	}
}

func (as *AuthService) CheckUserExists(email string, username string) (bool, error) {
	_, err := as.UserRepo.GetUserByEmail(email)
	if err == nil {
		return true, nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("failed to check if user exists by email (что-то не так): %w", err)
	}

	_, err = as.UserRepo.GetUserByUsername(username)
	if err == nil {
		return true, nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("failed to check if user exists by username: %w", err)
	}

	return false, nil
}

func (as *AuthService) Register(createUserDto dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	var userReponse *dto.UserResponseDTO
	user, err := models.NewUser(createUserDto.Email, createUserDto.Username, createUserDto.Password)
	if err != nil {
		return nil, err
	}
	if exists, err := as.CheckUserExists(createUserDto.Email, createUserDto.Username); err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("user already exists")
	}
	err = as.UserRepo.CreateUser(*user)
	if err != nil {
		return nil, err
	}
	userReponse = dto.UserToResponseDTO(user)
	return userReponse, nil
}

func (as *AuthService) Login(identifier string, password string) (*models.User, string, error) {
	user, err := as.UserRepo.GetUserByEmailOrUsername(identifier)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, "", fmt.Errorf("user not found %w", err)
		}
		return nil, "", fmt.Errorf("failed to get user by email/username %w", err)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		slog.Warn("Invalid password comparison", "identifier", identifier, "user.Password", user.Password, "password", password, "error", err)
		return nil, "", fmt.Errorf("invalid password %w", err)
	}
	token, err := utils.GenerateJWT(user.Id, as.JwtKey.SecretKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate jwtToken %w", err)
	}

	projects, err := as.ProjectRepo.GetAllProjects()
	if err != nil {
		return nil, "", fmt.Errorf("ошибка)")
	}
	as.WebSocketService.SendMessageToUser(user.Id, "projects", projects)
	return user, token, nil
}
