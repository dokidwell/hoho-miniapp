package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

// InitRedis 初始化Redis连接
func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
		PoolSize: 10,
	})

	// PING测试连接
	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Println("✓ Redis connection established")
	return nil
}

// CloseRedis 关闭Redis连接
func CloseRedis() {
	if RDB != nil {
		if err := RDB.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}
}
