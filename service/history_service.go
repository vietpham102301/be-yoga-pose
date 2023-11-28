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

func (hs *HistoryService) SaveHistoryLog(name string, path string, userID int, score int) error {
	return hs.historyRepo.SaveHistoryLog(name, path, userID, score)
}

func (hs *HistoryService) GetHistoryLog(pageSize int, pageNum int, userID int) (*[]models.History, error) {
	return hs.historyRepo.GetHistoryLog(pageSize, pageNum, userID)
}

func (hs *HistoryService) GetHistoryByID(id int) (*models.History, error) {
	return hs.historyRepo.GetHistoryByID(id)
}

func (hs *HistoryService) DeleteHistoryByID(id int, userID int) error {
	return hs.historyRepo.DeleteHistoryByID(id, userID)
}

func (hs *HistoryService) SavedHistory(id int, userID int) error {
	return hs.historyRepo.SavedHistory(id, userID)
}
