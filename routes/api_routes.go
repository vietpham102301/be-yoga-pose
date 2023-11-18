package routes

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yoga-pose-backend/handlers"
	"yoga-pose-backend/repository"
	"yoga-pose-backend/service"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	//r.Use(AuthMiddleware())
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	routers := r.Group("/api/v1/users")
	routers.GET("/:id", AuthMiddleware(), handlers.GetUserByIDHandler(userService))
	routers.POST("/register", handlers.RegisterUserHandler(userService))
	routers.POST("/login", handlers.LoginUserHandler(userService))
	routers.GET("/test", handlers.TestRunningPythonCode())

	routers.GET("/ws-video", handleClientConnection)

	return r
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("X-Access-Token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte("viet-secret-key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing error"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the token expiration
			exp := time.Unix(int64(claims["exp"].(float64)), 0)
			if time.Now().After(exp) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
				return
			}

			// Token is valid, proceed with the request
			c.Set("user", claims) // Store user information in the context
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
