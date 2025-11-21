package services

import (
	"time"

	"hoho-api/database"
	"hoho-api/models"
)

// GetAnnouncements 获取公告列表
func GetAnnouncements(announcementType string, page, pageSize int) ([]models.Announcement, int64, error) {
	var announcements []models.Announcement
	var total int64
	
	query := database.DB.Where("is_published = ?", true)
	
	if announcementType != "" {
		query = query.Where("type = ?", announcementType)
	}
	
	query.Model(&models.Announcement{}).Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("is_pinned DESC, published_at DESC").
		Offset(offset).Limit(pageSize).Find(&announcements).Error; err != nil {
		return nil, 0, err
	}
	
	return announcements, total, nil
}

// GetAnnouncementByID 根据ID获取公告
func GetAnnouncementByID(id uint) (*models.Announcement, error) {
	var announcement models.Announcement
	if err := database.DB.First(&announcement, id).Error; err != nil {
		return nil, err
	}
	return &announcement, nil
}

// CreateAnnouncement 创建公告
func CreateAnnouncement(title, content, announcementType, priority string, isPinned bool) (*models.Announcement, error) {
	announcement := &models.Announcement{
		Title:       title,
		Content:     content,
		Type:        announcementType,
		Priority:    priority,
		IsPinned:    isPinned,
		IsPublished: false,
	}
	
	if err := database.DB.Create(announcement).Error; err != nil {
		return nil, err
	}
	
	return announcement, nil
}

// PublishAnnouncement 发布公告
func PublishAnnouncement(id uint) error {
	now := time.Now()
	return database.DB.Model(&models.Announcement{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_published": true,
		"published_at": now,
	}).Error
}

// UnpublishAnnouncement 取消发布公告
func UnpublishAnnouncement(id uint) error {
	return database.DB.Model(&models.Announcement{}).Where("id = ?", id).Update("is_published", false).Error
}

// DeleteAnnouncement 删除公告
func DeleteAnnouncement(id uint) error {
	return database.DB.Delete(&models.Announcement{}, id).Error
}
