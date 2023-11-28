package models

type History struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Path      string `json:"path" `
	Score     int    `json:"score"`
	CreatedAt string `json:"created_at"`
	IsSaved   bool   `json:"is_saved"`
}
