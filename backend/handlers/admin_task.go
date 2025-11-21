package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/models"
	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetAllTasksAdmin 获取所有任务列表（管理员）
func GetAllTasksAdmin(c *gin.Context) {
	taskType := c.Query("type")
	enabled := c.Query("enabled")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	
	tasks, total, err := services.GetAllTasksAdmin(taskType, enabled, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"tasks":     tasks,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// CreateTask 创建任务
func CreateTask(c *gin.Context) {
	var task models.Task
	
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "任务创建成功",
		"task":    task,
	})
}

// UpdateTask 更新任务
func UpdateTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.UpdateTask(uint(taskID), updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "任务更新成功"})
}

// DeleteTask 删除任务
func DeleteTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "任务删除成功"})
}

// ToggleTask 启用/禁用任务
func ToggleTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		Enabled bool `json:"enabled"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.ToggleTask(uint(taskID), req.Enabled); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	status := "已禁用"
	if req.Enabled {
		status = "已启用"
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "任务" + status})
}
