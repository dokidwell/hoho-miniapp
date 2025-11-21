package handlers

import (
	"net/http"
	"strconv"

	"hoho-miniapp/backend/services"

	"github.com/gin-gonic/gin"
)

type AdminAnnouncementHandler struct {
	announcementService *services.AnnouncementService
}

func NewAdminAnnouncementHandler(announcementService *services.AnnouncementService) *AdminAnnouncementHandler {
	return &AdminAnnouncementHandler{
		announcementService: announcementService,
	}
}

// GetAnnouncementList 获取公告列表（管理员）
func (h *AdminAnnouncementHandler) GetAnnouncementList(c *gin.Context) {
	announcementType := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	announcements, total, err := h.announcementService.GetAnnouncements(announcementType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"announcements": announcements,
		"total":         total,
		"page":          page,
		"page_size":     pageSize,
	})
}

// CreateAnnouncement 创建公告
func (h *AdminAnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	var req struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		Type     string `json:"type" binding:"required"`
		Priority string `json:"priority"`
		IsPinned bool   `json:"is_pinned"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	announcement, err := h.announcementService.CreateAnnouncement(
		req.Title,
		req.Content,
		req.Type,
		req.Priority,
		req.IsPinned,
	)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":      "公告创建成功",
		"announcement": announcement,
	})
}

// UpdateAnnouncement 更新公告
func (h *AdminAnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Type     string `json:"type"`
		Priority string `json:"priority"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// TODO: 实现UpdateAnnouncement方法
	c.JSON(http.StatusOK, gin.H{"message": "公告更新成功", "id": announcementID})
}

// DeleteAnnouncement 删除公告
func (h *AdminAnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := h.announcementService.DeleteAnnouncement(uint(announcementID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告删除成功"})
}

// TogglePin 切换公告置顶状态
func (h *AdminAnnouncementHandler) TogglePin(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		IsPinned bool `json:"is_pinned"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// TODO: 实现TogglePin方法
	status := "已取消置顶"
	if req.IsPinned {
		status = "已置顶"
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告" + status, "id": announcementID})
}
