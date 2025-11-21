package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetAnnouncements 获取公告列表
func GetAnnouncements(c *gin.Context) {
	announcementType := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	announcements, total, err := services.GetAnnouncements(announcementType, page, pageSize)
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

// GetAnnouncementDetail 获取公告详情
func GetAnnouncementDetail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	announcement, err := services.GetAnnouncementByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "公告不存在"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"announcement": announcement})
}
