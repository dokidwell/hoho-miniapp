package models

import (
	"time"
)

// PlatformAccount 平台账户模型（阳光账户）
type PlatformAccount struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TotalBalance    string    `gorm:"type:decimal(20,8);not null;default:0.00000000" json:"total_balance"`
	CommissionIncome string   `gorm:"type:decimal(20,8);not null;default:0.00000000" json:"commission_income"`
	FeeIncome       string    `gorm:"type:decimal(20,8);not null;default:0.00000000" json:"fee_income"`
	TotalExpense    string    `gorm:"type:decimal(20,8);not null;default:0.00000000" json:"total_expense"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// PlatformTransaction 平台账户交易记录模型
type PlatformTransaction struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Type         string    `gorm:"type:enum('commission','fee','expense','adjustment');not null" json:"type"`
	Amount       string    `gorm:"type:decimal(20,8);not null" json:"amount"`
	BalanceAfter string    `gorm:"type:decimal(20,8);not null" json:"balance_after"`
	Description  string    `gorm:"size:255;not null" json:"description"`
	RelatedID    *uint     `json:"related_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// TableName 指定表名
func (PlatformAccount) TableName() string {
	return "platform_account"
}

// TableName 指定表名
func (PlatformTransaction) TableName() string {
	return "platform_transactions"
}
