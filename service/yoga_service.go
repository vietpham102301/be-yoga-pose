package service

import (
	"yoga-pose-backend/models"
	"yoga-pose-backend/repository"
)

type YogaService struct {
	yogaRepo *repository.YogaPoseRepository
}

func NewYogaService(yogaRepo *repository.YogaPoseRepository) *YogaService {
	return &YogaService{yogaRepo}
}

func (yg *YogaService) GetYogaPoseByName(yogaPoseName string) (*models.Yoga, error) {
	return yg.yogaRepo.GetPoseByName(yogaPoseName)
}
