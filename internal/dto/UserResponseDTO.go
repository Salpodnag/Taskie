package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	Id               uuid.UUID `json:"id"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	TimeRegistration time.Time `json:"timeRegistration"`
}
