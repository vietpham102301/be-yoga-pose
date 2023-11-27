package service

import (
	"yoga-pose-backend/models"
	"yoga-pose-backend/repository"
)

type HistoryService struct {
	historyRepo *repository.HistoryPoseRepository
}

func NewHistoryService(historyRepo *repository.HistoryPoseRepository) *HistoryService {
	return &HistoryService{historyRepo}
}

func (hs *HistoryService) SaveHistoryLog(name string, path string, userID int) error {
	return hs.historyRepo.SaveHistoryLog(name, path, userID)
}

func (hs *HistoryService) GetHistoryLog(pageSize int, pageNum int, userID int) (*[]models.History, error) {
	return hs.historyRepo.GetHistoryLog(pageSize, pageNum, userID)
}
