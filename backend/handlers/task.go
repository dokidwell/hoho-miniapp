package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetTasks 获取任务列表
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	tasks, err := services.GetUserTasks(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// ClaimTaskReward 领取任务奖励
func ClaimTaskReward(c *gin.Context) {
	userID, _ := c.Get("user_id")
	completionID, _ := strconv.ParseUint(c.Param("completion_id"), 10, 32)
	
	if err := services.ClaimTaskReward(userID.(uint), uint(completionID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "奖励领取成功"})
}

// DailySignIn 每日签到
func DailySignIn(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	if err := services.DailySignIn(userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "签到成功"})
}
