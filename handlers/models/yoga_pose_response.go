package models

import (
	"yoga-pose-backend/models"
)

type YogaPoseResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path" `
}

func (c *YogaPoseResponse) ToYogaPoseResponse(yoga *models.Yoga) *YogaPoseResponse {
	return &YogaPoseResponse{
		ID:   yoga.ID,
		Name: yoga.Name,
		Path: yoga.Path,
	}
}
