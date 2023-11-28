package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"yoga-pose-backend/service"
)

func GetYogaPoseByName(yogaService *service.YogaService) func(*gin.Context) {
	return func(c *gin.Context) {
		yogaPoseName := c.Query("poseName")
		if yogaPoseName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "poseName parameter is required"})
			return
		}
		yogaPose, err := yogaService.GetYogaPoseByName(yogaPoseName)

		imagePath := yogaPose.Path

		if imagePath == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image path not found"})
			return
		}

		imageData, err := os.ReadFile(imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
			return
		}

		c.Header("Content-Type", "image/png")

		c.Header("Content-Disposition", "attachment; filename="+yogaPoseName+".png")

		c.Data(http.StatusOK, "image/png", imageData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve yoga pose"})
			return
		}
	}
}
