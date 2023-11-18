package repository

import (
	"database/sql"
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

func (ur *UserRepository) AuthenticateUser(user *models.UserLoginRequest) (*models.User, error) {
	// Query the database to get the user by username
	query := "SELECT * FROM users WHERE username = ?"
	userData := &models.User{}
	err := ur.db.QueryRow(query, user.Username).Scan(&userData.ID, &userData.Username, &userData.Email, &userData.PasswordHash, &userData.CreatedAt, &userData.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
