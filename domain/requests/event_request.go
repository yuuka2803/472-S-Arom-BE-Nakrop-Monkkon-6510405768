package requests

import "github.com/google/uuid"

type CreateEventRequest struct {
	Title       	string    `json:"title" db:"Event_Title"`
	Description 	string    `json:"description" db:"Event_Description"`
	Start       	string    `json:"start" db:"Event_Start"`
	End         	string    `json:"end" db:"Event_End"`
	Tag         	uuid.UUID    `json:"tag" db:"Event_Tag"`
	Notification    bool      `json:"notification" db:"Event_Email"`
	ReminderAt        string    `json:"reminder" db:"Event_Reminder"`
	UserId      	uuid.UUID `json:"user_id" db:"User_Id"`
}

type UpdateEventRequest struct {
	Completed bool `json:"completed" db:"Event_Complete"`
}
