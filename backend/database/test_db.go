package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitTestDatabase 初始化测试数据库（使用SQLite）
func InitTestDatabase() error {
	// 使用SQLite内存数据库进行测试
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 连接SQLite数据库
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to test database: %w", err)
	}

	log.Println("✓ Test database connection established (SQLite in-memory)")
	
	// 初始化测试数据
	return InitTestData()
}
