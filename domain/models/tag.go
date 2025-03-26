package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	Id          uuid.UUID `json:"id" db:"Tag_ID"`
	Name       string    `json:"name" db:"Tag_Name"`
	UserId      uuid.UUID `json:"user_id" db:"User_Id"`
}
