package models

import (
	"time"
)

// Announcement 公告模型
type Announcement struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:100;not null" json:"title"`
	Content     string     `gorm:"type:text;not null" json:"content"`
	Type        string     `gorm:"type:enum('system','rule','event','maintenance');default:'system'" json:"type"`
	Priority    string     `gorm:"type:enum('low','normal','high','urgent');default:'normal'" json:"priority"`
	IsPinned    bool       `gorm:"default:false" json:"is_pinned"`
	IsPublished bool       `gorm:"default:false" json:"is_published"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Notification 通知模型
type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Type      string    `gorm:"type:enum('offer','trade','task','announcement','system');not null" json:"type"`
	Title     string    `gorm:"size:100;not null" json:"title"`
	Content   string    `gorm:"size:255;not null" json:"content"`
	RelatedID *uint     `json:"related_id"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	
	// 关联
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName 指定表名
func (Announcement) TableName() string {
	return "announcements"
}

// TableName 指定表名
func (Notification) TableName() string {
	return "notifications"
}
