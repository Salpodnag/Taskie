package services

import (
	"Taskie/cfg"
	"Taskie/internal/models"
	"Taskie/internal/repositories"
	"Taskie/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	JwtKey   cfg.JWT
	UserRepo repositories.UserRepository
}

func NewAuthService(JWT cfg.JWT, ur repositories.UserRepository) *AuthService {
	return &AuthService{
		JwtKey:   JWT,
		UserRepo: ur,
	}
}

func (as *AuthService) CheckUserExists(email string, username string) (bool, error) {
	_, err := as.UserRepo.GetUserByEmail(email)
	if err == nil {
		return true, nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("failed to check if user exists by email: %w", err)
	}

	_, err = as.UserRepo.GetUserByUsername(username)
	if err == nil {
		return true, nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("failed to check if user exists by username: %w", err)
	}

	return false, nil
}

func (as *AuthService) Register(email string, username string, password string) error {
	var user models.User
	hashedPassword := utils.HashFromPassword(password)
	user.Email = email
	user.Username = username
	user.Password = string(hashedPassword)
	user.TimeRegistration = time.Now()
	if exists, err := as.CheckUserExists(email, username); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("user already exists")
	}
	err := as.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
