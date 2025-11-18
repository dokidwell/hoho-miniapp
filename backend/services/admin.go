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

// GetUsersPaginated 获取分页用户列表
func (s *AdminService) GetUsersPaginated(page, pageSize int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	query := database.DB.Model(&models.User{})

	if search != "" {
		// 模糊搜索UID或手机号
		query = query.Where("uid LIKE ? OR phone LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetUserDetail 获取用户详情，包括积分
func (s *AdminService) GetUserDetail(userID uint64) (*models.User, *models.UserPoint, error) {
	var user models.User
	result := database.DB.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, nil, result.Error
	}

	var points models.UserPoint
	result = database.DB.Where("user_id = ?", userID).First(&points)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 理论上不会发生，但为了健壮性，返回一个空的积分结构
		return &user, &models.UserPoint{UserID: userID}, nil
	}
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return &user, &points, nil
}

// UpdateUserStatus 更新用户状态
func (s *AdminService) UpdateUserStatus(userID uint64, status string) error {
	return database.DB.Model(&models.User{}).Where("id = ?", userID).Update("status", status).Error
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
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
