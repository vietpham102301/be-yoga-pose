package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}

// HashPassword Function to hash a user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword Function to verify a user's password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IssueJWTToken Function to issue a JWT token
func IssueJWTToken(user User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	})
	tokenString, err := token.SignedString([]byte("viet-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
