package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hoho-miniapp/backend/services"
)

type AdminTaskHandler struct {
	taskService *services.TaskService
}

func NewAdminTaskHandler(taskService *services.TaskService) *AdminTaskHandler {
	return &AdminTaskHandler{
		taskService: taskService,
	}
}

// GetTaskList 获取任务列表（管理员）
func (h *AdminTaskHandler) GetTaskList(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// CreateTask 创建任务
func (h *AdminTaskHandler) CreateTask(c *gin.Context) {
	// TODO: 实现创建任务
	c.JSON(http.StatusOK, gin.H{"message": "任务创建成功"})
}

// UpdateTask 更新任务
func (h *AdminTaskHandler) UpdateTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	// TODO: 实现更新任务
	c.JSON(http.StatusOK, gin.H{"message": "任务更新成功", "id": taskID})
}

// DeleteTask 删除任务
func (h *AdminTaskHandler) DeleteTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	// TODO: 实现删除任务
	c.JSON(http.StatusOK, gin.H{"message": "任务删除成功", "id": taskID})
}

// ToggleTask 启用/禁用任务
func (h *AdminTaskHandler) ToggleTask(c *gin.Context) {
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		IsEnabled bool `json:"is_enabled"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// TODO: 实现启用/禁用任务
	status := "已禁用"
	if req.IsEnabled {
		status = "已启用"
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "任务" + status, "id": taskID})
}
