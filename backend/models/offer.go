package models

import (
	"time"
)

// Offer 出价模型（心愿单功能）
type Offer struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	BuyerID            uint       `gorm:"not null;index" json:"buyer_id"`
	ArtworkInstanceID  uint       `gorm:"not null;index" json:"artwork_instance_id"`
	Price              string     `gorm:"type:decimal(20,8);not null" json:"price"`
	Status             string     `gorm:"type:enum('pending','accepted','rejected','cancelled','expired');default:'pending'" json:"status"`
	ExpiresAt          *time.Time `json:"expires_at"`
	RespondedAt        *time.Time `json:"responded_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	
	// 关联
	Buyer           *User             `gorm:"foreignKey:BuyerID" json:"buyer,omitempty"`
	ArtworkInstance *ArtworkInstance  `gorm:"foreignKey:ArtworkInstanceID" json:"artwork_instance,omitempty"`
}

// TableName 指定表名
func (Offer) TableName() string {
	return "offers"
}
