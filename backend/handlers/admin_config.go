package handlers

import (
	"net/http"

	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetSystemConfigs 获取所有系统配置
func GetSystemConfigs(c *gin.Context) {
	configs, err := services.GetAllSystemConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"configs": configs})
}

// UpdateSystemConfig 更新系统配置
func UpdateSystemConfig(c *gin.Context) {
	key := c.Param("key")
	
	var req struct {
		Value       string `json:"value" binding:"required"`
		Description string `json:"description"`
		Reason      string `json:"reason"` // 修改原因（用于透明公示）
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	adminID, _ := c.Get("admin_id")
	
	if err := services.UpdateSystemConfig(key, req.Value, req.Description, req.Reason, adminID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "配置更新成功"})
}

// GetConfigHistory 获取配置修改历史
func GetConfigHistory(c *gin.Context) {
	key := c.Query("key")
	
	history, err := services.GetConfigHistory(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"history": history})
}
