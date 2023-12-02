package repository

import (
	"database/sql"
	"errors"
	"fmt"
	models2 "yoga-pose-backend/handlers/models"
	"yoga-pose-backend/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	// Query the database to get the user by ID
	query := "SELECT * FROM users WHERE id = ?"
	user := &models.User{}
	err := ur.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) RegisterUser(user *models.User) (*models.User, error) {
	// Insert the user into the database
	query := "INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)"
	_, err := ur.db.Exec(query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) AuthenticateUser(user *models2.UserLoginRequest) (*models.User, error) {
	// Query the database to get the user by username
	query := "SELECT * FROM users WHERE username = ?"
	userData := &models.User{}
	err := ur.db.QueryRow(query, user.Username).Scan(&userData.ID, &userData.Username, &userData.Email, &userData.PasswordHash, &userData.CreatedAt, &userData.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (ur *UserRepository) ResetPassword(emailOrUsername string, passwordHashed string) (*models.User, error) {
	query := "UPDATE users SET password_hash = ? WHERE email = ? OR username = ?"
	_, err := ur.db.Exec(query, passwordHashed, emailOrUsername, emailOrUsername)
	if err != nil {
		return nil, err
	}

	selectQuery := "SELECT * FROM users WHERE email = ? OR username = ?"
	row := ur.db.QueryRow(selectQuery, emailOrUsername, emailOrUsername)
	var user models.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with email/username %s not found", emailOrUsername)
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdatePassword(userID int, newPasswordHashed string, oldPassword string) error {
	//check if the old password is correct
	query := "SELECT * FROM users WHERE id = ?"
	user := &models.User{}
	err := ur.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	err = models.VerifyPassword(user.PasswordHash, oldPassword)
	if err != nil {
		return err
	}

	//update the new password
	query = "UPDATE users SET password_hash = ? WHERE id = ?"
	_, err = ur.db.Exec(query, newPasswordHashed, userID)
	if err != nil {
		return err
	}
	return nil
}
