package models

type Yoga struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path" `
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
