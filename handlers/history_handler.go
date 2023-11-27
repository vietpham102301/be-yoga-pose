package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yoga-pose-backend/handlers/models"
	"yoga-pose-backend/service"
)

func GetHistoryHandler(service *service.HistoryService) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		pageSize := c.Query("pageSize")
		pageNum := c.Query("pageNum")

		if pageSize == "" {
			pageSize = "10"
		}
		if pageNum == "" {
			pageNum = "1"
		}
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}
		pageNumInt, err := strconv.Atoi(pageNum)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		history, err := service.GetHistoryLog(pageSizeInt, pageNumInt, userIDInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		c.JSON(200, history)
	}

}

func SaveHistoryHandler(service *service.HistoryService) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		historyRequest := &models.HistorySaveRequest{}
		if err := c.ShouldBindJSON(historyRequest); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		err = service.SaveHistoryLog(historyRequest.Name, historyRequest.Path, userIDInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		c.JSON(200, gin.H{"message": "Successfully saved history log"})
	}
}
