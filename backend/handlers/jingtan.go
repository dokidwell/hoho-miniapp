package handlers

import (
	"hoho-miniapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JingtanHandler 定义鲸探API相关的HTTP处理函数
type JingtanHandler struct {
	JingtanService *services.JingtanService
}

// NewJingtanHandler 创建一个新的JingtanHandler实例
func NewJingtanHandler(jingtanService *services.JingtanService) *JingtanHandler {
	return &JingtanHandler{JingtanService: jingtanService}
}

// BindAccount 处理绑定鲸探账户请求
func (h *JingtanHandler) BindAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		JingtanAccountID string `json:"jingtan_account_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	if err := h.JingtanService.BindAccount(userID.(uint64), req.JingtanAccountID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "绑定失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "绑定成功",
	})
}

// UnbindAccount 处理解绑鲸探账户请求
func (h *JingtanHandler) UnbindAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if err := h.JingtanService.UnbindAccount(userID.(uint64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "解绑失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "解绑成功",
	})
}

// SyncAssets 处理同步鲸探资产请求
func (h *JingtanHandler) SyncAssets(c *gin.Context) {
	userID, _ := c.Get("user_id")

	assets, err := h.JingtanService.SyncAssets(userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "同步失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "同步成功",
		"data": assets,
	})
}

// GetAssets 处理获取鲸探资产列表请求
func (h *JingtanHandler) GetAssets(c *gin.Context) {
	userID, _ := c.Get("user_id")

	assets, err := h.JingtanService.GetUserJingtanAssets(userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取资产列表失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": assets,
	})
}
