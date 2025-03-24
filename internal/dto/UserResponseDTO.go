package dto

import (
	"Taskie/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	Id               uuid.UUID `json:"id"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	TimeRegistration time.Time `json:"timeRegistration"`
}

func UserToResponseDTO(user *models.User) *UserResponseDTO {
	return &UserResponseDTO{
		Id:               user.Id,
		Email:            user.Email,
		Username:         user.Username,
		TimeRegistration: user.TimeRegistration,
	}
}
