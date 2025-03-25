package requests

import "github.com/google/uuid"

type CreateTagRequest struct {
	Name   string    `json:"name" db:"Tag_Name"`
	UserID uuid.UUID `json:"user_id"`
}

type UpdateTagRequest struct {
	Name     string    `json:"name" db:"Tag_Name"`   
}