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
	CreatedAt   time.Time `json:"createdAt"`
}

func NewProject(name string, description string, color string, privacy string) (*Project, error) {
	if err := validateProject(name); err != nil {
		return nil, err
	}
	return &Project{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Color:       color,
		Privacy:     privacy,
		CreatedAt:   time.Now(),
	}, nil
}

func validateProject(name string) error {
	if name == "" {
		return fmt.Errorf("empty project name")
	}
	return nil
}
