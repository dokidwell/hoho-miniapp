package models

import (
	"time"

	"gorm.io/gorm"
)

// Collection 藏品集合/系列
type Collection struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(500)" json:"description"`
	CoverImage  string `gorm:"type:varchar(500)" json:"cover_image"`
	Status      string `gorm:"type:enum('active', 'inactive');default:'active'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Asset 藏品（SKU）
type Asset struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey" json:"id"`
	CollectionID uint64 `gorm:"index;not null" json:"collection_id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(500)" json:"description"`
	MediaURL    string `gorm:"type:varchar(500);not null" json:"media_url"`
	MediaType   string `gorm:"type:enum('image', 'video', 'audio');not null" json:"media_type"`
	TotalSupply int    `gorm:"not null" json:"total_supply"` // 总发行量
	MintedCount int    `gorm:"default:0" json:"minted_count"` // 已铸造数量
	CreatorID   uint64 `gorm:"index;not null" json:"creator_id"` // 创作者ID
	Status      string `gorm:"type:enum('pending_review', 'approved', 'rejected', 'active', 'inactive');default:'pending_review'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AssetInstance 藏品实例（具体编号）
type AssetInstance struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey" json:"id"`
	AssetID     uint64 `gorm:"index;not null" json:"asset_id"`
	InstanceNo  int    `gorm:"not null" json:"instance_no"` // 实例编号（#1, #2, #3...）
	OwnerID     uint64 `gorm:"index;not null" json:"owner_id"` // 当前持有者ID
	TokenID     string `gorm:"type:varchar(255);uniqueIndex;not null" json:"token_id"` // 唯一标识符，模拟链上TokenID
	Status      string `gorm:"type:enum('in_wallet', 'on_sale', 'pending_trade', 'burned');default:'in_wallet'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// JingtanAsset 鲸探资产映射表
type JingtanAsset struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey" json:"id"`
	UserID          uint64 `gorm:"index;not null" json:"user_id"`
	JingtanAssetID  string `gorm:"type:varchar(255);uniqueIndex;not null" json:"jingtan_asset_id"` // 鲸探平台的资产ID
	Name            string `gorm:"type:varchar(100)" json:"name"`
	ImageURL        string `gorm:"type:varchar(500)" json:"image_url"`
	Status          string `gorm:"type:enum('active', 'inactive');default:'active'" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
