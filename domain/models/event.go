package models
import (
	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id" db:"Event_Id"`
	Name        string `json:"name" db:"Event_Name"`
	Description string `json:"description" db:"Event_Description"`
}
