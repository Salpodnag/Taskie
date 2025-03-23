package models

import (
	"Taskie/internal/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID `json:"id"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	TimeRegistration time.Time `json:"timeRegistration"`
}

func NewUser(email string, username string, password string) (*User, error) {
	if err := validateUser(email, username); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}
	return &User{
		Id:               uuid.New(),
		Email:            email,
		Username:         username,
		Password:         string(utils.HashFromPassword(password)),
		TimeRegistration: time.Now(),
	}, nil
}

func validateUser(email string, username string) error {
	if email == "" {
		return fmt.Errorf("empty user email: %w")
	}
	if username == "" {
		return fmt.Errorf("empty user username: %w")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("empty user password: %w")
	}
	return nil
}
