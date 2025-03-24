package requests

import (
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Email		 string `json:"email" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SendEmailRequest struct {
	ID		uuid.UUID `json:"userID" validate:"required"`
	Email   string    `json:"email" validate:"required"`
}
