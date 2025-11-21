package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/models"
	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// CreateAnnouncement 创建公告
func CreateAnnouncement(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	
	var announcement models.Announcement
	if err := c.ShouldBindJSON(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	announcement.CreatedBy = adminID.(uint)
	
	if err := services.CreateAnnouncement(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":      "公告创建成功",
		"announcement": announcement,
	})
}

// UpdateAnnouncement 更新公告
func UpdateAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.UpdateAnnouncement(uint(announcementID), updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告更新成功"})
}

// PublishAnnouncement 发布公告
func PublishAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.PublishAnnouncement(uint(announcementID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告已发布"})
}

// DeleteAnnouncement 删除公告
func DeleteAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.DeleteAnnouncement(uint(announcementID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告删除成功"})
}

// PinAnnouncement 置顶公告
func PinAnnouncement(c *gin.Context) {
	announcementID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		Pinned bool `json:"pinned"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.PinAnnouncement(uint(announcementID), req.Pinned); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	status := "已取消置顶"
	if req.Pinned {
		status = "已置顶"
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "公告" + status})
}

// GetAllAnnouncementsAdmin 获取所有公告（管理员）
func GetAllAnnouncementsAdmin(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	announcements, total, err := services.GetAllAnnouncementsAdmin(status, page, pageSize)
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
