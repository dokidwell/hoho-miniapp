package services

import (
	"errors"
	"fmt"
	"time"

	"hoho-api/database"
	"hoho-api/models"
	
	"github.com/shopspring/decimal"
)

// GetAllTasks 获取所有启用的任务
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("is_enabled = ?", true).Order("sort_order ASC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetUserTasks 获取用户的任务列表（包含完成状态）
func GetUserTasks(userID uint) ([]map[string]interface{}, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, err
	}
	
	result := make([]map[string]interface{}, 0)
	
	for _, task := range tasks {
		taskData := map[string]interface{}{
			"id":            task.ID,
			"name":          task.Name,
			"description":   task.Description,
			"type":          task.Type,
			"reward_points": task.RewardPoints,
			"completed":     false,
			"claimed":       false,
		}
		
		// 检查任务完成状态
		if task.Type == "daily" {
			// 每日任务：检查今天是否完成
			var completion models.UserTaskCompletion
			today := time.Now().Format("2006-01-02")
			err := database.DB.Where("user_id = ? AND task_id = ? AND DATE(completed_at) = ?", 
				userID, task.ID, today).First(&completion).Error
			
			if err == nil {
				taskData["completed"] = true
				taskData["claimed"] = completion.ClaimedAt != nil
				taskData["completion_id"] = completion.ID
			}
		} else {
			// 一次性任务或成就任务：检查是否完成过
			var completion models.UserTaskCompletion
			err := database.DB.Where("user_id = ? AND task_id = ?", userID, task.ID).
				Order("completed_at DESC").First(&completion).Error
			
			if err == nil {
				taskData["completed"] = true
				taskData["claimed"] = completion.ClaimedAt != nil
				taskData["completion_id"] = completion.ID
			}
		}
		
		result = append(result, taskData)
	}
	
	return result, nil
}

// CompleteTask 完成任务
func CompleteTask(userID, taskID uint) error {
	// 获取任务信息
	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		return err
	}
	
	if !task.IsEnabled {
		return errors.New("任务已禁用")
	}
	
	// 检查是否已完成
	if task.Type == "daily" {
		// 每日任务：检查今天是否已完成
		var count int64
		today := time.Now().Format("2006-01-02")
		database.DB.Model(&models.UserTaskCompletion{}).
			Where("user_id = ? AND task_id = ? AND DATE(completed_at) = ?", userID, taskID, today).
			Count(&count)
		
		if count > 0 {
			return errors.New("今日已完成该任务")
		}
	} else {
		// 一次性任务：检查是否已完成过
		var count int64
		database.DB.Model(&models.UserTaskCompletion{}).
			Where("user_id = ? AND task_id = ?", userID, taskID).
			Count(&count)
		
		if count > 0 {
			return errors.New("该任务已完成")
		}
	}
	
	// 记录任务完成
	completion := &models.UserTaskCompletion{
		UserID:       userID,
		TaskID:       taskID,
		CompletedAt:  time.Now(),
		RewardPoints: task.RewardPoints,
	}
	
	if err := database.DB.Create(completion).Error; err != nil {
		return err
	}
	
	// 发送通知
	notification := &models.Notification{
		UserID:  userID,
		Type:    "task",
		Title:   "任务完成",
		Content: fmt.Sprintf("恭喜！您完成了任务《%s》，可领取 %s 积分", task.Name, task.RewardPoints),
	}
	database.DB.Create(notification)
	
	return nil
}

// ClaimTaskReward 领取任务奖励
func ClaimTaskReward(userID, completionID uint) error {
	// 获取任务完成记录
	var completion models.UserTaskCompletion
	if err := database.DB.Preload("Task").First(&completion, completionID).Error; err != nil {
		return err
	}
	
	// 检查是否属于该用户
	if completion.UserID != userID {
		return errors.New("无权领取该任务奖励")
	}
	
	// 检查是否已领取
	if completion.ClaimedAt != nil {
		return errors.New("奖励已领取")
	}
	
	// 发放积分
	points, _ := decimal.NewFromString(completion.RewardPoints)
	if err := AddPoints(userID, points, "task", fmt.Sprintf("任务奖励：%s", completion.Task.Name), &completionID); err != nil {
		return err
	}
	
	// 更新领取状态
	now := time.Now()
	if err := database.DB.Model(&completion).Update("claimed_at", now).Error; err != nil {
		return err
	}
	
	return nil
}

// DailySignIn 每日签到
func DailySignIn(userID uint) error {
	// 获取签到任务
	var task models.Task
	if err := database.DB.Where("condition_type = ? AND is_enabled = ?", "daily_signin", true).First(&task).Error; err != nil {
		return errors.New("签到任务不存在")
	}
	
	return CompleteTask(userID, task.ID)
}

// CheckAndCompleteFirstPurchase 检查并完成首次购买任务
func CheckAndCompleteFirstPurchase(userID uint) {
	// 检查是否是首次购买
	var count int64
	database.DB.Model(&models.Trade{}).Where("buyer_id = ?", userID).Count(&count)
	
	if count == 1 {
		// 首次购买，完成任务
		var task models.Task
		if err := database.DB.Where("condition_type = ? AND is_enabled = ?", "first_purchase", true).First(&task).Error; err == nil {
			CompleteTask(userID, task.ID)
		}
	}
}

// CheckAndCompleteFirstSale 检查并完成首次出售任务
func CheckAndCompleteFirstSale(userID uint) {
	// 检查是否是首次出售
	var count int64
	database.DB.Model(&models.Trade{}).Where("seller_id = ?", userID).Count(&count)
	
	if count == 1 {
		// 首次出售，完成任务
		var task models.Task
		if err := database.DB.Where("condition_type = ? AND is_enabled = ?", "first_sale", true).First(&task).Error; err == nil {
			CompleteTask(userID, task.ID)
		}
	}
}
