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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "X-Access-Token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour}))

	//r.Use(AuthMiddleware())
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	yogaService := service.NewYogaService(repository.NewYogaRepository(db))
	historyService := service.NewHistoryService(repository.NewHistoryRepository(db))

	routers := r.Group("/api/v1/users")
	routers.GET("/:id", AuthMiddleware(), handlers.GetUserByIDHandler(userService))
	routers.POST("/register", handlers.RegisterUserHandler(userService))
	routers.POST("/login", handlers.LoginUserHandler(userService))
	routers.GET("/test", handlers.TestRunningPythonCode())
	routers.GET("/ws-video", handleClientConnection)

	routers = r.Group("/api/v1/yoga")
	routers.GET("/pose", handlers.GetYogaPoseByName(yogaService))

	routers = r.Group("/api/v1/history")
	routers.GET("/pose", AuthMiddleware(), handlers.GetHistoryHandler(historyService))
	routers.POST("/pose", AuthMiddleware(), handlers.SaveHistoryHandler(historyService))
	routers.GET("/pose/:id", handlers.GetHistoryImageHandler(historyService))
	routers.DELETE("/pose/:id", AuthMiddleware(), handlers.DeleteHistoryHandler(historyService))
	routers.PUT("/pose/:id", AuthMiddleware(), handlers.SavedHistoryHandler(historyService))

	return r
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-Access-Token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := time.Unix(int64(claims["exp"].(float64)), 0)
			if time.Now().After(exp) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
				return
			}

			c.Set("user", claims)
			c.Set("userID", claims["user_id"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
