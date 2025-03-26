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
		TimeRegistration: time.Now().UTC().Add(time.Hour * 3),
	}, nil
}

func validateUser(email string, username string) error {
	if email == "" {
		return fmt.Errorf("empty user email")
	}
	if username == "" {
		return fmt.Errorf("empty user username")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("empty user password")
	}
	return nil
}
