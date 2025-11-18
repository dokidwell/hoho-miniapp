package services

import (
	"errors"
	"fmt"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
	"hoho-miniapp/backend/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// UserService 定义用户服务接口
type UserService struct{}

// NewUserService 创建一个新的UserService实例
func NewUserService() *UserService {
	return &UserService{}
}

// Register 注册新用户
func (s *UserService) Register(phone, password string) (*models.User, string, error) {
	// 1. 检查用户是否已存在
	var existingUser models.User
	result := database.DB.Where("phone = ?", phone).First(&existingUser)
	if result.Error == nil {
		return nil, "", errors.New("手机号已被注册")
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, "", result.Error
	}

	// 2. 哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", errors.New("密码加密失败")
	}

	// 3. 创建用户
	user := models.User{
		Phone:        phone,
		PasswordHash: hashedPassword,
		Nickname:     fmt.Sprintf("HOHO用户%s", phone[7:]), // 默认昵称
		UID:          utils.GenerateUID(), // 生成UID
		Status:       "active",
	}

	// 4. 开启事务
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// 4.1. 创建用户
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// 4.2. 初始化用户积分
		userPoint := models.UserPoint{
			UserID: user.ID,
			Balance: decimal.Zero,
			Frozen: decimal.Zero,
			TotalEarned: decimal.Zero,
			TotalSpent: decimal.Zero,
		}
		if err := tx.Create(&userPoint).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, "", err
	}

	// 5. 生成JWT Token
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("Token生成失败")
	}

	return &user, token, nil
}

// Login 用户登录
func (s *UserService) Login(phone, password string) (*models.User, string, error) {
	var user models.User
	result := database.DB.Where("phone = ?", phone).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, "", errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, "", result.Error
	}

	// 检查密码
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, "", errors.New("用户已被禁用")
	}

	// 生成JWT Token
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("Token生成失败")
	}

	return &user, token, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint64) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	return &user, result.Error
}

// GenerateToken 生成JWT Token
func (s *UserService) GenerateToken(userID uint64) (string, error) {
	// 从环境变量获取密钥和过期时间
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET未设置")
	}
	
	// 默认过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
