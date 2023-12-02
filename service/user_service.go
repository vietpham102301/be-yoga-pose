package service

import (
	models2 "yoga-pose-backend/handlers/models"
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

func (us *UserService) AuthenticateUser(user *models2.UserLoginRequest) (*models.User, error) {
	return us.userRepo.AuthenticateUser(user)
}

func (us *UserService) ResetPassword(emailOrUsername string, passwordHashed string) (*models.User, error) {
	return us.userRepo.ResetPassword(emailOrUsername, passwordHashed)
}

func (us *UserService) UpdatePassword(userID int, newPasswordHashed string, oldPassword string) error {
	return us.userRepo.UpdatePassword(userID, newPasswordHashed, oldPassword)
}
