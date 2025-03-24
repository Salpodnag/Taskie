package dto

import "fmt"

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateUserDTO) ValidateCreateUser() error {
	if dto.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if dto.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if dto.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	return nil
}
