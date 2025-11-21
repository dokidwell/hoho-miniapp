package database

import (
	"hoho-miniapp/backend/models"
	"log"
)

// InitTestData 初始化测试数据
func InitTestData() error {
	// 自动迁移测试所需的表（使用简化模型）
	err := DB.AutoMigrate(
		&models.TestUser{},
		&models.SystemConfig{},
		&models.Task{},
		&models.Announcement{},
		&models.PlatformAccount{},
	)
	
	if err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return err
	}
	
	log.Println("数据库表创建成功")
	
	// 创建测试管理员
	admin := &models.Admin{
		Username: "admin",
		Role:     "super_admin",
	}
	DB.Create(admin)
	
	// 创建系统配置
	configs := []models.SystemConfig{
		{Key: "default_commission_rate", Value: "70", Description: "创作者默认分成比例"},
		{Key: "platform_commission_rate", Value: "30", Description: "平台分成比例"},
		{Key: "trade_fee_rate", Value: "2.5", Description: "交易手续费率"},
	}
	for _, config := range configs {
		DB.Create(&config)
	}
	
	// 创建测试任务
	tasks := []models.Task{
		{
			Name:          "每日签到",
			Description:   "每天签到可获得积分",
			Type:          "daily",
			ConditionType: "daily_signin",
			RewardPoints:  "10",
			IsEnabled:     true,
			SortOrder:     1,
		},
		{
			Name:          "首次购买",
			Description:   "完成首次购买可获得奖励",
			Type:          "once",
			ConditionType: "first_purchase",
			RewardPoints:  "100",
			IsEnabled:     true,
			SortOrder:     2,
		},
	}
	for _, task := range tasks {
		DB.Create(&task)
	}
	
	// 创建测试公告
	announcements := []models.Announcement{
		{
			Title:    "欢迎来到HOHO平台",
			Content:  "这是一个测试公告",
			Type:     "system",
			Priority: "normal",
			IsPinned: true,
		},
	}
	for _, announcement := range announcements {
		DB.Create(&announcement)
	}
	
	// 创建平台账户
	platformAccount := &models.PlatformAccount{
		TotalBalance:     "0",
		CommissionIncome: "0",
		FeeIncome:        "0",
		TotalExpense:     "0",
	}
	DB.Create(platformAccount)
	
	log.Println("测试数据初始化完成")
	return nil
}
