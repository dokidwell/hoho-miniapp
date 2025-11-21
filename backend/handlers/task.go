package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hoho-miniapp/backend/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// GetTaskList 获取任务列表
func (h *TaskHandler) GetTaskList(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	tasks, err := h.taskService.GetUserTasks(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// CompleteTask 完成任务
func (h *TaskHandler) CompleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	completion, err := h.taskService.CompleteTask(userID.(uint), uint(taskID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "任务完成",
		"completion": completion,
	})
}

// ClaimReward 领取任务奖励
func (h *TaskHandler) ClaimReward(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := h.taskService.ClaimTaskReward(userID.(uint), uint(taskID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "奖励领取成功"})
}
