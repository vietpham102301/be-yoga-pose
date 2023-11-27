package models

type SavedPose struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Path      string `json:"path" `
	CreatedAt string `json:"created_at"`
}
