package models

import (
	"time"
)

// Creation 创作模型
type Creation struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"not null;index" json:"user_id"`
	Title          string     `gorm:"size:100;not null" json:"title"`
	Description    string     `gorm:"type:text" json:"description"`
	MediaURL       string     `gorm:"size:255;not null" json:"media_url"`
	ThumbnailURL   string     `gorm:"size:255" json:"thumbnail_url"`
	TotalSupply    uint       `gorm:"not null" json:"total_supply"`
	Price          string     `gorm:"type:decimal(20,8);not null" json:"price"`
	CommissionRate string     `gorm:"type:decimal(5,2);not null;default:40.00" json:"commission_rate"` // 平台分成比例（%）
	Status         string     `gorm:"type:enum('pending','approved','rejected','published');default:'pending'" json:"status"`
	RejectReason   string     `gorm:"size:255" json:"reject_reason"`
	ReviewerID     *uint      `json:"reviewer_id"`
	ReviewedAt     *time.Time `json:"reviewed_at"`
	PublishedAt    *time.Time `json:"published_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	
	// 关联
	User     *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Reviewer *User `gorm:"foreignKey:ReviewerID" json:"reviewer,omitempty"`
}

// TableName 指定表名
func (Creation) TableName() string {
	return "creations"
}
