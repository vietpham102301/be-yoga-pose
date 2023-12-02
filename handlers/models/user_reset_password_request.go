package models

type UserResetPasswordRequest struct {
	EmailOrUsername string `json:"email_or_username"`
}
