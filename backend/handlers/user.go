package handlers

import (
	"hoho-miniapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 定义用户相关的HTTP处理函数
type UserHandler struct {
	UserService *services.UserService
}

// NewUserHandler 创建一个新的UserHandler实例
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// RegisterRequest 定义注册请求体
type RegisterRequest struct {
	Phone           string `json:"phone" binding:"required,len=11"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

// Register 处理用户注册请求
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	user, token, err := h.UserService.Register(req.Phone, req.Password)
	if err != nil {
		// 检查是否是手机号已注册的错误
		if err.Error() == "手机号已被注册" {
			c.JSON(http.StatusConflict, gin.H{"code": 409, "message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "注册失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
		"data": gin.H{
			"id":       user.ID,
			"uid":      user.UID,
			"phone":    user.Phone,
			"nickname": user.Nickname,
			"token":    token,
		},
	})
}

// LoginRequest 定义登录请求体
type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required"`
}

// Login 处理用户登录请求
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	user, token, err := h.UserService.Login(req.Phone, req.Password)
	if err != nil {
		// 检查特定错误
		if err.Error() == "用户不存在" {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": err.Error()})
			return
		}
		if err.Error() == "密码错误" {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1002, "message": err.Error()})
			return
		}
		if err.Error() == "用户已被禁用" {
			c.JSON(http.StatusForbidden, gin.H{"code": 1003, "message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "登录失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data": gin.H{
			"id":         user.ID,
			"uid":        user.UID,
			"phone":      user.Phone,
			"nickname":   user.Nickname,
			"avatar_url": user.AvatarURL,
			"token":      token,
		},
	})
}

// GetProfile 处理获取用户资料请求
func (h *UserHandler) GetProfile(c *gin.Context) {
	// 从上下文中获取用户ID，这是通过JWT中间件设置的
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	user, err := h.UserService.GetUserByID(userID.(uint64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}

// UpdateProfileRequest 定义更新资料请求体
type UpdateProfileRequest struct {
	Nickname  *string `json:"nickname"`
	AvatarURL *string `json:"avatar_url"`
}

// UpdateProfile 处理更新用户资料请求
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	updates := make(map[string]interface{})

	if req.Nickname != nil {
		updates["nickname"] = *req.Nickname
	}
	if req.AvatarURL != nil {
		updates["avatar_url"] = *req.AvatarURL
	}

	user, err := h.UserService.UpdateProfile(userID.(uint64), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新成功",
		"data":    user,
	})
}

// VerifyIdentityRequest 定义实名认证请求体
type VerifyIdentityRequest struct {
	RealName string `json:"real_name" binding:"required"`
	IDNumber string `json:"id_number" binding:"required"`
}

// VerifyIdentity 处理实名认证请求
func (h *UserHandler) VerifyIdentity(c *gin.Context) {
	var req VerifyIdentityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	user, err := h.UserService.VerifyIdentity(userID.(uint64), req.RealName, req.IDNumber)
	if err != nil {
		if err.Error() == "用户已完成实名认证" {
			c.JSON(http.StatusConflict, gin.H{"code": 409, "message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "实名认证失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "实名认证成功",
		"data": gin.H{
			"id":                user.ID,
			"real_name":         user.RealName,
			"identity_verified": user.IdentityVerified,
		},
	})
}

// GetPoints 处理获取用户积分请求
func (h *UserHandler) GetPoints(c *gin.Context) {
	userID, _ := c.Get("user_id")

	points, err := h.UserService.GetUserPoints(userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取积分失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    points,
	})
}
