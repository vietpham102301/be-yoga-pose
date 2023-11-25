package models

import (
	"yoga-pose-backend/models"
)

type UserLoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *UserLoginResponse) ToUserLoginResponse(user *models.User) *UserLoginResponse {
	return &UserLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
