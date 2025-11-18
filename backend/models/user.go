package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// User 对应数据库中的 users 表
type User struct {
	gorm.Model
	ID               uint64    `gorm:"primaryKey" json:"id"`
	UID              string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"uid"`
	Phone            string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	PasswordHash     string    `gorm:"type:varchar(255);not null" json:"-"`
	Nickname         string    `gorm:"type:varchar(100)" json:"nickname"`
	AvatarURL        string    `gorm:"type:varchar(500)" json:"avatar_url"`
	RealName         string    `gorm:"type:varchar(100)" json:"real_name"`
	IDNumber         string    `gorm:"type:varchar(50)" json:"-"`
	IdentityVerified bool      `gorm:"default:false" json:"identity_verified"`
	Status           string    `gorm:"type:enum('active', 'suspended', 'banned');default:'active'" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	// 关联关系
	Points UserPoint `gorm:"foreignKey:UserID"`
}

// UserPoint 对应数据库中的 user_points 表
type UserPoint struct {
	gorm.Model
	ID          uint64          `gorm:"primaryKey" json:"id"`
	UserID      uint64          `gorm:"uniqueIndex;not null" json:"user_id"`
	Balance     decimal.Decimal `gorm:"type:decimal(30,8);default:0" json:"balance"`
	Frozen      decimal.Decimal `gorm:"type:decimal(30,8);default:0" json:"frozen"`
	TotalEarned decimal.Decimal `gorm:"type:decimal(30,8);default:0" json:"total_earned"`
	TotalSpent  decimal.Decimal `gorm:"type:decimal(30,8);default:0" json:"total_spent"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// PointTransaction 对应数据库中的 point_transactions 表
type PointTransaction struct {
	gorm.Model
	ID          uint64          `gorm:"primaryKey" json:"id"`
	UserID      uint64          `gorm:"index;not null" json:"user_id"`
	Type        string          `gorm:"type:enum('earn', 'spend', 'adjust', 'freeze', 'unfreeze')" json:"type"`
	Amount      decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"amount"`
	Description string          `gorm:"type:varchar(255)" json:"description"`
	RelatedID   uint64          `gorm:"index" json:"related_id"`
	RelatedType string          `gorm:"type:varchar(50);index" json:"related_type"`
	CreatedAt   time.Time       `json:"created_at"`
}

// ThirdPartyAccount 对应数据库中的 third_party_accounts 表
type ThirdPartyAccount struct {
	gorm.Model
	ID               uint64    `gorm:"primaryKey" json:"id"`
	UserID           uint64    `gorm:"index;not null" json:"user_id"`
	Platform         string    `gorm:"type:varchar(50);not null" json:"platform"`
	PlatformUID      string    `gorm:"type:varchar(255);not null" json:"platform_uid"`
	PlatformUsername string    `gorm:"type:varchar(255)" json:"platform_username"`
	AccessToken      string    `gorm:"type:varchar(500)" json:"-"`
	RefreshToken     string    `gorm:"type:varchar(500)" json:"-"`
	TokenExpiresAt   time.Time `json:"token_expires_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// TableName 为 ThirdPartyAccount 指定表名
func (ThirdPartyAccount) TableName() string {
	return "third_party_accounts"
}
