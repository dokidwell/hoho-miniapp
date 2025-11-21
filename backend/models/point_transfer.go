package models

import (
	"time"
)

// PointTransfer 积分转账模型（预留功能）
type PointTransfer struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FromUserID uint      `gorm:"not null;index" json:"from_user_id"`
	ToUserID   uint      `gorm:"not null;index" json:"to_user_id"`
	Amount     string    `gorm:"type:decimal(20,8);not null" json:"amount"`
	Fee        string    `gorm:"type:decimal(20,8);not null;default:0.00000000" json:"fee"`
	Message    string    `gorm:"size:255" json:"message"`
	Status     string    `gorm:"type:enum('pending','completed','failed');default:'pending'" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	
	// 关联
	FromUser *User `gorm:"foreignKey:FromUserID" json:"from_user,omitempty"`
	ToUser   *User `gorm:"foreignKey:ToUserID" json:"to_user,omitempty"`
}

// TableName 指定表名
func (PointTransfer) TableName() string {
	return "point_transfers"
}
