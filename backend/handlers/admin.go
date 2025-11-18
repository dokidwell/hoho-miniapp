package handlers

import (
	"hoho-miniapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminHandler 定义管理员相关的HTTP处理函数
type AdminHandler struct {
	AdminService *services.AdminService
}

// NewAdminHandler 创建一个新的AdminHandler实例
func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{AdminService: adminService}
}

// LoginRequest 定义登录请求体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 处理管理员登录请求
func (h *AdminHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	admin, token, err := h.AdminService.Login(req.Username, req.Password)
	if err != nil {
		// 检查特定错误
		if err.Error() == "管理员不存在" || err.Error() == "密码错误" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "用户名或密码错误"})
			return
		}
		if err.Error() == "管理员已被禁用" {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "登录失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "登录成功",
		"data": gin.H{
			"id": admin.ID,
			"username": admin.Username,
			"role": admin.Role,
			"token": token,
		},
	})
}

// GetProfile 处理获取管理员资料请求
func (h *AdminHandler) GetProfile(c *gin.Context) {
	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	admin, err := h.AdminService.GetAdminByID(adminID.(uint64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "管理员不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": admin,
	})
}
