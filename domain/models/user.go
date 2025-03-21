package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"User_Id"`
	Username     string    `json:"username" db:"Username"`
	Password     string    `json:"password" db:"Password"`
	Email		 string    `json:"email" db:"Email"`
	ProfileImage string    `json:"profile_image" db:"Profile_Image"`
}
