package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"net/http"
	"net/smtp"
	"os/exec"
	"strconv"
	"strings"
	"time"
	models2 "yoga-pose-backend/handlers/models"
	"yoga-pose-backend/models"
	"yoga-pose-backend/service"
)

var emailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Your Password have been reset</title>
</head>
<body>
	<h3>{{.GeneralUser}}</h3>
	<p>{{.NewPassword}}</p>
</body>
</html>
`

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

func ResetPasswordHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		var resetPasswordRequest models2.UserResetPasswordRequest
		if err := c.ShouldBindJSON(&resetPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		randomPassword := randomPasswordGenerator(8)
		passwordHashed, err := models.HashPassword(randomPassword)

		user, err := userService.ResetPassword(resetPasswordRequest.EmailOrUsername, passwordHashed)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = sendResetPasswordToEmail(user.Email, randomPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Your new password has been sent to your email"})
	}
}

func randomPasswordGenerator(length int) string {
	rand.NewSource(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// sendResetPasswordToEmail sends a reset password email to the user.
func sendResetPasswordToEmail(userEmail string, newPassword string) error {
	// Email template data
	data := struct {
		GeneralUser string
		NewPassword string
	}{
		GeneralUser: "Dear my awesome user!",
		NewPassword: "Your new password is: " + newPassword + ". Please change it after you login.",
	}

	// Parse the email template
	tmpl, err := template.New("resetPassword").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("error parsing email template: %v", err)
	}

	// Execute the template and store the output in a buffer
	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, data)
	if err != nil {
		return fmt.Errorf("error executing email template: %v", err)
	}

	// Set up authentication information
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "peterpham102301@gmail.com"
	smtpPassword := "ptrk thug viaw etfq"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Set up the email message
	to := []string{userEmail}
	msg := []byte("To: " + userEmail + "\r\n" +
		"Subject: Reset Your Password\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		emailBody.String())

	// Connect to the server and send the email with TLS
	err = smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, smtpUsername, to, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}

func UpdatePasswordHandler(userService *service.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		var updatePasswordRequest models2.UserUpdatePasswordRequest
		if err := c.ShouldBindJSON(&updatePasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		userIDFromToken := c.MustGet("userID").(float64)
		userIDFromTokenInt := int(userIDFromToken)
		newPasswordHashed, err := models.HashPassword(updatePasswordRequest.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = userService.UpdatePassword(userIDFromTokenInt, newPasswordHashed, updatePasswordRequest.OldPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
	}
}
