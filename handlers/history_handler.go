package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"yoga-pose-backend/config"
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
			pageNum = "0"
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

		userIDFromToken := c.MustGet("userID").(float64)
		userIDFromTokenInt := int(userIDFromToken)

		if userIDInt != userIDFromTokenInt {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		history, err := service.GetHistoryLog(pageSizeInt, pageNumInt, userIDInt)
		for i := 0; i < len(*history); i++ {
			(*history)[i].Path = config.HostURL + "/api/v1/history/pose/" + strconv.FormatInt((*history)[i].ID, 10)
		}
		res := []models.HistoryResponse{}
		for _, h := range *history {
			temp := models.HistoryResponse{}
			res = append(res, *temp.ToHistoryResponse(&h))
		}

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		c.JSON(200, res)
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

		userIDFromToken := c.MustGet("userID").(float64)
		userIDFromTokenInt := int(userIDFromToken)

		if userIDInt != userIDFromTokenInt {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		err = service.SaveHistoryLog(historyRequest.Name, historyRequest.Path, userIDInt, historyRequest.Score)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		c.JSON(200, gin.H{"message": "Successfully saved history log"})
	}
}

func GetHistoryImageHandler(service *service.HistoryService) func(*gin.Context) {
	return func(c *gin.Context) {
		historyID := c.Param("id")
		historyIDInt, err := strconv.Atoi(historyID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		history, err := service.GetHistoryByID(historyIDInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}
		imageData, err := os.ReadFile(history.Path)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve history log"})
			return
		}

		c.Header("Content-Type", "image/png")

		c.Header("Content-Disposition", "attachment; filename=history.png")

		c.Data(200, "image/png", imageData)
	}
}

func DeleteHistoryHandler(service *service.HistoryService) func(*gin.Context) {
	return func(c *gin.Context) {
		historyID := c.Param("id")
		historyIDInt, err := strconv.Atoi(historyID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete history log"})
			return
		}

		userIDFromToken := c.MustGet("userID").(float64)
		userIDFromTokenInt := int(userIDFromToken)

		err = service.DeleteHistoryByID(historyIDInt, userIDFromTokenInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete history log"})
			return
		}

		c.JSON(200, gin.H{"message": "Successfully deleted history log"})
	}
}

func SavedHistoryHandler(service *service.HistoryService) func(*gin.Context) {
	return func(c *gin.Context) {
		historyID := c.Param("id")
		historyIDInt, err := strconv.Atoi(historyID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		userIDFromToken := c.MustGet("userID").(float64)
		userIDFromTokenInt := int(userIDFromToken)

		err = service.SavedHistory(historyIDInt, userIDFromTokenInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save history log"})
			return
		}

		c.JSON(200, gin.H{"message": "Successfully saved history log"})
	}
}
