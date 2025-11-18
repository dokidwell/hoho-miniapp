package handlers

import (
	"hoho-miniapp/backend/services"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminHandler 定义管理员相关的HTTP处理函数
type AdminHandler struct {
	AdminService *services.AdminService
	AssetService *services.AssetService
}

// NewAdminHandler 创建一个新的AdminHandler实例
func NewAdminHandler(adminService *services.AdminService, assetService *services.AssetService) *AdminHandler {
	return &AdminHandler{
		AdminService: adminService,
		AssetService: assetService,
	}
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

// ListUsersPage 处理用户列表页面请求
func (h *AdminHandler) ListUsersPage(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := 20
	search := c.DefaultQuery("search", "")

	pageNum := 1
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageNum = p
	}

	users, total, err := h.AdminService.GetUsersPaginated(pageNum, pageSize, search)
	if err != nil {
		c.String(http.StatusInternalServerError, "获取用户列表失败: "+err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	
	// 简化分页导航，只显示当前页前后2页
	startPage := pageNum - 2
	if startPage < 1 {
		startPage = 1
	}
	endPage := pageNum + 2
	if endPage > totalPages {
		endPage = totalPages
	}

	var pages []int
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}

	c.HTML(http.StatusOK, "admin_users.html", gin.H{
		"Title": "用户管理",
		"ActiveMenu": "users",
		"Users": users,
		"Total": total,
		"Page": pageNum,
		"PageSize": pageSize,
		"TotalPages": totalPages,
		"Search": search,
		"Pages": pages,
	})
}

// UpdateUserStatus 处理更新用户状态请求
func (h *AdminHandler) UpdateUserStatus(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "用户ID格式错误"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=active suspended banned"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	if err := h.AdminService.UpdateUserStatus(userID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新用户状态失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})
}

// ListAssetReviewPage 处理藏品审核列表页面请求
func (h *AdminHandler) ListAssetReviewPage(c *gin.Context) {
	assets, err := h.AssetService.GetPendingReviewAssets()
	if err != nil {
		c.String(http.StatusInternalServerError, "获取待审核藏品列表失败: "+err.Error())
		return
	}

	c.HTML(http.StatusOK, "admin_asset_review.html", gin.H{
		"Title": "铸造审核",
		"ActiveMenu": "review",
		"Assets": assets,
	})
}

// ReviewAsset 处理藏品审核请求
func (h *AdminHandler) ReviewAsset(c *gin.Context) {
	assetIDStr := c.Param("id")
	assetID, err := strconv.ParseUint(assetIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "藏品ID格式错误"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=approved rejected"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	_, err = h.AssetService.ReviewMintRequest(assetID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "审核操作失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "审核操作成功"})
}
