package service

import (
	"yoga-pose-backend/models"
	"yoga-pose-backend/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (us *UserService) GetUserByID(userID int) (*models.User, error) {
	return us.userRepo.GetUserByID(userID)
}

func (us *UserService) RegisterUser(user *models.User) (*models.User, error) {
	return us.userRepo.RegisterUser(user)
}

func (us *UserService) AuthenticateUser(user *models.UserLoginRequest) (*models.User, error) {
	return us.userRepo.AuthenticateUser(user)
}
