package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminConfigHandler struct{}

func NewAdminConfigHandler() *AdminConfigHandler {
	return &AdminConfigHandler{}
}

// GetConfig 获取所有系统配置
func (h *AdminConfigHandler) GetConfig(c *gin.Context) {
	// TODO: 实现获取系统配置
	c.JSON(http.StatusOK, gin.H{
		"configs": map[string]interface{}{
			"creator_commission_rate":  "70",
			"platform_commission_rate": "30",
			"trade_fee_rate":           "2.5",
			"enable_creation":          true,
			"enable_task":              true,
			"enable_third_party":       true,
		},
	})
}

// UpdateConfig 更新系统配置
func (h *AdminConfigHandler) UpdateConfig(c *gin.Context) {
	var req map[string]interface{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// TODO: 实现更新系统配置
	c.JSON(http.StatusOK, gin.H{"message": "配置更新成功"})
}
