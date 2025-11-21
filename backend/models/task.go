package models

import (
	"time"
)

// Task 任务模型
type Task struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"size:50;not null" json:"name"`
	Description    string    `gorm:"size:255" json:"description"`
	Type           string    `gorm:"type:enum('daily','once','achievement');not null" json:"type"`
	RewardPoints   string    `gorm:"type:decimal(20,8);not null" json:"reward_points"`
	ConditionType  string    `gorm:"size:50;not null" json:"condition_type"`
	ConditionValue string    `gorm:"size:255" json:"condition_value"`
	IsEnabled      bool      `gorm:"default:true" json:"is_enabled"`
	SortOrder      int       `gorm:"default:0" json:"sort_order"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UserTaskCompletion 用户任务完成记录模型
type UserTaskCompletion struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	UserID       uint       `gorm:"not null;index" json:"user_id"`
	TaskID       uint       `gorm:"not null;index" json:"task_id"`
	CompletedAt  time.Time  `json:"completed_at"`
	ClaimedAt    *time.Time `json:"claimed_at"`
	RewardPoints string     `gorm:"type:decimal(20,8);not null" json:"reward_points"`
	
	// 关联
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Task *Task `gorm:"foreignKey:TaskID" json:"task,omitempty"`
}

// TableName 指定表名
func (Task) TableName() string {
	return "tasks"
}

// TableName 指定表名
func (UserTaskCompletion) TableName() string {
	return "user_task_completions"
}
