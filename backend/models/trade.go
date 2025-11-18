package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Listing 挂售单
type Listing struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey" json:"id"`
	AssetInstanceID uint64 `gorm:"index;not null" json:"asset_instance_id"`
	SellerID        uint64 `gorm:"index;not null" json:"seller_id"`
	Price           decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"price"`
	Status          string `gorm:"type:enum('active', 'sold', 'canceled');default:'active'" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Trade 交易记录
type Trade struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey" json:"id"`
	ListingID       uint64 `gorm:"index;not null" json:"listing_id"`
	AssetInstanceID uint64 `gorm:"index;not null" json:"asset_instance_id"`
	BuyerID         uint64 `gorm:"index;not null" json:"buyer_id"`
	SellerID        uint64 `gorm:"index;not null" json:"seller_id"`
	Price           decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"price"`
	PlatformFee     decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"platform_fee"` // 平台手续费（2.5%）
	CreatorRoyalty  decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"creator_royalty"` // 创作者版税（2.5%）
	SellerReceived  decimal.Decimal `gorm:"type:decimal(30,8);not null" json:"seller_received"` // 卖家实际收到
	Status          string `gorm:"type:enum('pending', 'completed', 'failed', 'canceled');default:'pending'" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// CommunityEvent 社区事件（用于透明公示）
type CommunityEvent struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey" json:"id"`
	EventType   string `gorm:"type:varchar(50);index;not null" json:"event_type"` // 事件类型：mint, airdrop, trade, burn, etc.
	UserID      uint64 `gorm:"index" json:"user_id"`
	Description string `gorm:"type:varchar(500)" json:"description"`
	RelatedID   uint64 `gorm:"index" json:"related_id"` // 关联ID（如Trade ID, Asset ID）
	RelatedType string `gorm:"type:varchar(50)" json:"related_type"` // 关联类型
	CreatedAt   time.Time `json:"created_at"`
}
