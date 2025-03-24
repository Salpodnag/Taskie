package dto

import "fmt"

type CreateProjectDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Privacy     string `json:"privacy"`
}

func (dto *CreateProjectDTO) Validate() error {
	if dto.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	if dto.Description == "" {
		return fmt.Errorf("project description cannot be empty")
	}
	if dto.Color == "" {
		return fmt.Errorf("project color cannot be empty")
	}
	if dto.Privacy == "" {
		return fmt.Errorf("project privacy cannot be empty")
	}
	return nil
}
