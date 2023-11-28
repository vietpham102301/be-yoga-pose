package models

import (
	"yoga-pose-backend/config"
	models2 "yoga-pose-backend/models"
)

type HistoryResponse struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	SamplePath string `json:"sample_path"`
	Score      int    `json:"score"`
	CreatedAt  string `json:"created_at"`
	IsSaved    bool   `json:"is_saved"`
}

func (h *HistoryResponse) ToHistoryResponse(history *models2.History) *HistoryResponse {
	samplePath := config.HostURL + "/api/v1/yoga/pose?poseName=" + history.Name
	return &HistoryResponse{
		ID:         history.ID,
		UserID:     history.UserID,
		Name:       history.Name,
		Path:       history.Path,
		SamplePath: samplePath,
		Score:      history.Score,
		CreatedAt:  history.CreatedAt,
		IsSaved:    history.IsSaved,
	}
}
