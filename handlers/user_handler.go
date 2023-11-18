package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
	"time"
	"yoga-pose-backend/models"
	"yoga-pose-backend/service"
)

func GetUserByIDHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		userIDParam := c.Param("id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func RegisterUserHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse the request body to get user registration data
		var registrationRequest models.UserRegistrationRequest
		if err := c.ShouldBindJSON(&registrationRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}
		// parse request body to get user registration data

		//hash the password
		hashedPassword, err := models.HashPassword(registrationRequest.Password)
		userData := models.User{
			Username:     registrationRequest.Username,
			Email:        registrationRequest.Email,
			PasswordHash: hashedPassword,
		}
		// Call the service to register the user
		user, err := userService.RegisterUser(&userData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If registration is successful, return a response
		c.JSON(http.StatusCreated, user)
	}
}

func LoginUserHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginRequest models.UserLoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// Authenticate the user using the service
		user, err := userService.AuthenticateUser(&loginRequest)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			return
		}

		err = models.VerifyPassword(user.PasswordHash, loginRequest.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			return
		}

		// Issue a JWT token upon successful authentication
		token, err := models.IssueJWTToken(*user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue JWT token"})
			return
		}
		user.PasswordHash = ""
		// Return the user data as JSON response
		c.Header("X-Access-Token", token)
		c.JSON(http.StatusOK, user)

		// Set the JWT token in the response header

		fmt.Println("Response Headers:", c.Writer.Header())
	}
}

func TestRunningPythonCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		imagePath := "F:\\saved_frames\\8775_1699378658.png"
		cmd := exec.Command("cmd", "/C", "G:\\smt\\run_python.bat", imagePath)
		elapsed := time.Since(start)
		fmt.Printf("page took %s", elapsed)
		stdout, err := cmd.CombinedOutput()

		if err != nil {
			println(err.Error())
			return
		}
		println(string(stdout))
		c.JSON(http.StatusOK, gin.H{"data": string(stdout)})

	}
}
