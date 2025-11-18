package utils

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 检查密码哈希是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateUID 生成一个唯一的UID
func GenerateUID() string {
	// 简单实现：时间戳 + 随机数
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("U%d%04d", time.Now().Unix()%1000000, rand.Intn(10000))
}
