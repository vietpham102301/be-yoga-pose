package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yoga-pose-backend/handlers/models"
	"yoga-pose-backend/service"
)

func SavePoseHandler(savePoseService *service.SavePoseService) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		savePoseRequest := &models.SavePoseRequest{}
		if err := c.ShouldBindJSON(savePoseRequest); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save pose"})
		}
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save pose"})
			return
		}

		err = savePoseService.SavePose(userIDInt, savePoseRequest.Name, savePoseRequest.Path)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save pose"})
			return
		}

		c.JSON(200, gin.H{"success": "Pose saved"})
	}

}

func GetSavedPosesHandler(savePoseService *service.SavePoseService) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		pageSize := c.Query("pageSize")
		pageNum := c.Query("pageNum")
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get saved poses"})
			return
		}
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get saved poses"})
			return
		}

		pageNumInt, err := strconv.Atoi(pageNum)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get saved poses"})
			return
		}
		savedPoses, err := savePoseService.GetSavedPoses(pageSizeInt, pageNumInt, userIDInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get saved poses"})
			return
		}

		c.JSON(200, savedPoses)
	}
}
