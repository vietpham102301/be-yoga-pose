package service

import (
	"yoga-pose-backend/models"
	"yoga-pose-backend/repository"
)

type SavePoseService struct {
	savePoseRepo *repository.SavePoseRepository
}

func NewSavePoseService(savePoseRepo *repository.SavePoseRepository) *SavePoseService {
	return &SavePoseService{savePoseRepo}
}

func (ss *SavePoseService) SavePose(userID int, name string, path string) error {
	return ss.savePoseRepo.SavePose(userID, name, path)
}

func (ss *SavePoseService) GetSavedPoses(pageSize int, pageNum int, userID int) (*[]models.SavedPose, error) {
	return ss.savePoseRepo.GetSavedPoses(pageSize, pageNum, userID)
}
