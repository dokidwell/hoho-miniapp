package services

import (
	"errors"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
	"hoho-miniapp/backend/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// AdminService 定义管理员服务接口
type AdminService struct{}

// NewAdminService 创建一个新的AdminService实例
func NewAdminService() *AdminService {
	return &AdminService{}
}

// Login 管理员登录
func (s *AdminService) Login(username, password string) (*models.Admin, string, error) {
	var admin models.Admin
	result := database.DB.Where("username = ?", username).First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, "", errors.New("管理员不存在")
	}
	if result.Error != nil {
		return nil, "", result.Error
	}

	// 检查密码
	if !utils.CheckPasswordHash(password, admin.PasswordHash) {
		return nil, "", errors.New("密码错误")
	}

	// 检查管理员状态
	if admin.Status != "active" {
		return nil, "", errors.New("管理员已被禁用")
	}

	// 更新最后登录时间
	database.DB.Model(&admin).Update("last_login_at", time.Now())

	// 生成JWT Token
	token, err := s.GenerateToken(admin.ID)
	if err != nil {
		return nil, "", errors.New("Token生成失败")
	}

	return &admin, token, nil
}

// GetAdminByID 根据ID获取管理员
func (s *AdminService) GetAdminByID(id uint64) (*models.Admin, error) {
	var admin models.Admin
	result := database.DB.First(&admin, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("管理员不存在")
	}
	return &admin, result.Error
}

// GenerateToken 生成JWT Token
func (s *AdminService) GenerateToken(adminID uint64) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET未设置")
	}
	
	// 管理员Token有效期可以设置长一些，例如7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"admin_id": adminID,
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
