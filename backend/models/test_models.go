package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// TestUser 测试用的简化用户模型
type TestUser struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	UID              string         `gorm:"uniqueIndex;size:20;not null" json:"uid"`
	Phone            string         `gorm:"uniqueIndex;size:20;not null" json:"phone"`
	PasswordHash     string         `gorm:"size:255;not null" json:"-"`
	Nickname         string         `gorm:"size:100" json:"nickname"`
	AvatarURL        string         `gorm:"size:500" json:"avatar_url"`
	RealName         string         `gorm:"size:100" json:"real_name"`
	IDNumber         string         `gorm:"size:50" json:"id_number"`
	IdentityVerified bool           `gorm:"default:false" json:"identity_verified"`
	Status           string         `gorm:"size:20;default:active" json:"status"` // active, suspended, banned
	Points           decimal.Decimal `gorm:"type:decimal(20,8);default:0" json:"points"`
}

func (TestUser) TableName() string {
	return "users"
}
