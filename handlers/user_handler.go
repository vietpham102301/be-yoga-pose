package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
	"time"
	models2 "yoga-pose-backend/handlers/models"
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
		var registrationRequest models2.UserRegistrationRequest
		if err := c.ShouldBindJSON(&registrationRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		hashedPassword, err := models.HashPassword(registrationRequest.Password)
		userData := models.User{
			Username:     registrationRequest.Username,
			Email:        registrationRequest.Email,
			PasswordHash: hashedPassword,
		}
		user, err := userService.RegisterUser(&userData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func LoginUserHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginRequest models2.UserLoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

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

		token, err := models.IssueJWTToken(*user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue JWT token"})
			return
		}
		res := models2.UserLoginResponse{}

		c.Header("X-Access-Token", token)
		c.JSON(http.StatusOK, res.ToUserLoginResponse(user))

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
