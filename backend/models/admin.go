package models

import (
	"time"

	"gorm.io/gorm"
)

// Admin 对应数据库中的 admins 表
type Admin struct {
	gorm.Model
	ID           uint64    `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Role         string    `gorm:"size:20;default:'admin'" json:"role"`
	Status       string    `gorm:"size:20;default:'active'" json:"status"`
	LastLoginAt  time.Time `json:"last_login_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName 为 Admin 指定表名
func (Admin) TableName() string {
	return "admins"
}
