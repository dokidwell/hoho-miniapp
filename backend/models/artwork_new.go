package models

import (
	"time"
)

// Artwork 作品模型（统一术语：asset → artwork）
type Artwork struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Title        string     `gorm:"size:100;not null" json:"title"`
	Description  string     `gorm:"type:text" json:"description"`
	CreatorName  string     `gorm:"size:50" json:"creator_name"`
	MediaType    string     `gorm:"type:enum('image','video','audio','3d');default:'image'" json:"media_type"`
	MediaURL     string     `gorm:"size:255;not null" json:"media_url"`
	ThumbnailURL string     `gorm:"size:255" json:"thumbnail_url"`
	TotalSupply  uint       `gorm:"not null" json:"total_supply"`
	MintedCount  uint       `gorm:"default:0" json:"minted_count"`
	Price        string     `gorm:"type:decimal(20,8);not null" json:"price"` // 使用string存储decimal
	Source       string     `gorm:"type:enum('jingtan','platform','community','waveup');default:'platform'" json:"source"`
	Series       string     `gorm:"size:50" json:"series"`
	ReleaseDate  *time.Time `json:"release_date"`
	Status       string     `gorm:"type:enum('draft','pending','active','sold_out','archived');default:'draft'" json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// ArtworkInstance 作品实例模型
type ArtworkInstance struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ArtworkID    uint      `gorm:"not null;index" json:"artwork_id"`
	OwnerID      uint      `gorm:"not null;index" json:"owner_id"`
	SerialNumber string    `gorm:"size:50;uniqueIndex;not null" json:"serial_number"`
	MintedAt     time.Time `json:"minted_at"`
	Status       string    `gorm:"type:enum('owned','listed','locked');default:'owned'" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	
	// 关联
	Artwork *Artwork `gorm:"foreignKey:ArtworkID" json:"artwork,omitempty"`
	Owner   *User    `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
}

// TableName 指定表名
func (Artwork) TableName() string {
	return "artworks"
}

// TableName 指定表名
func (ArtworkInstance) TableName() string {
	return "artwork_instances"
}
