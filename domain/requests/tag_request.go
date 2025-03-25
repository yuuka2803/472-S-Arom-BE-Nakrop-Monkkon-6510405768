package requests

type CreateTagRequest struct{
	Name      string    `json:"name" db:"Tag_Name"`
	UserID      string   `json:"user_id"`
}

type UpdateTagRequest struct {
	Completed bool `json:"completed" db:"Tag_Complete"`
}

