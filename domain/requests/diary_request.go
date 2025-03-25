package requests

type CreateDiaryRequest struct {
	Date        string   `json:"date" `
	Mood        string   `json:"mood"`
	Emotions    []string `json:"emotions"`
	Description string   `json:"description"`
	UserID      string   `json:"user_id"`
	Images      []string `json:"images"`
}

type UpdateDiaryRequest struct {
	Date        string   `json:"date" `
	UserID      string   `json:"user_id"`
	Mood        string   `json:"mood"`
	Emotions    []string `json:"emotions"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	ImagesURL   []string `json:"images_url"`
}
