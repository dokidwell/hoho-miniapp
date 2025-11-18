package config

import (
	"os"
	"strconv"

	"github.com/shopspring/decimal"
)

// Config 系统配置
type Config struct {
	// 交易相关配置
	PlatformFeeRate  decimal.Decimal // 平台手续费比例（默认2.5%）
	CreatorRoyaltyRate decimal.Decimal // 创作者版税比例（默认2.5%）
	
	// 积分相关配置
	InitialPoints decimal.Decimal // 新用户初始积分（默认100.00000000）
	
	// 系统配置
	DecimalPrecision int32 // 积分精度（默认8位小数）
}

var AppConfig *Config

// InitConfig 初始化配置
func InitConfig() {
	AppConfig = &Config{
		PlatformFeeRate:  getDecimalEnv("PLATFORM_FEE_RATE", "0.025"),      // 2.5%
		CreatorRoyaltyRate: getDecimalEnv("CREATOR_ROYALTY_RATE", "0.025"), // 2.5%
		InitialPoints:    getDecimalEnv("INITIAL_POINTS", "100.00000000"),
		DecimalPrecision: 8,
	}
}

// getDecimalEnv 从环境变量获取Decimal值，如果不存在则使用默认值
func getDecimalEnv(key, defaultValue string) decimal.Decimal {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	
	result, err := decimal.NewFromString(value)
	if err != nil {
		// 如果解析失败，使用默认值
		result, _ = decimal.NewFromString(defaultValue)
	}
	
	return result
}

// getIntEnv 从环境变量获取整数值，如果不存在则使用默认值
func getIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	
	return result
}
