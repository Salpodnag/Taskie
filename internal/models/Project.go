package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Privacy     string    `json:"privacy"`
	Owner       User      `json:"user"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewProject(name string, owner User, description string, color string, privacy string) (*Project, error) {
	if err := validateProject(name, owner); err != nil {
		return nil, err
	}
	return &Project{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Color:       color,
		Privacy:     privacy,
		Owner:       owner,
		CreatedAt:   time.Now(),
	}, nil
}

func validateProject(name string, owner User) error {
	if name == "" {
		return fmt.Errorf("empty project name")
	}
	if err := validateUser(owner.Email, owner.Username); err != nil {
		return fmt.Errorf("empty owner: %w", err)
	}
	return nil
}
