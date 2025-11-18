package services

import (
	"errors"
	"fmt"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"

	"gorm.io/gorm"
)

// JingtanService 定义鲸探API对接服务接口
type JingtanService struct{}

// NewJingtanService 创建一个新的JingtanService实例
func NewJingtanService() *JingtanService {
	return &JingtanService{}
}

// BindAccount 绑定鲸探账户（模拟）
func (s *JingtanService) BindAccount(userID uint64, jingtanAccountID string) error {
	// 检查是否已绑定
	var existingAccount models.ThirdPartyAccount
	result := database.DB.Where("user_id = ? AND platform = ?", userID, "jingtan").First(&existingAccount)
	if result.Error == nil {
		return errors.New("已绑定鲸探账户")
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// TODO: 实际项目中需要调用鲸探API进行OAuth授权
	// 这里模拟绑定成功

	account := models.ThirdPartyAccount{
		UserID:      userID,
		Platform:    "jingtan",
		PlatformUID: jingtanAccountID,
		Status:      "active",
	}

	if err := database.DB.Create(&account).Error; err != nil {
		return err
	}

	return nil
}

// UnbindAccount 解绑鲸探账户
func (s *JingtanService) UnbindAccount(userID uint64) error {
	result := database.DB.Where("user_id = ? AND platform = ?", userID, "jingtan").Delete(&models.ThirdPartyAccount{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未绑定鲸探账户")
	}
	return nil
}

// SyncAssets 同步鲸探资产（模拟）
func (s *JingtanService) SyncAssets(userID uint64) ([]models.JingtanAsset, error) {
	// 检查是否已绑定鲸探账户
	var account models.ThirdPartyAccount
	result := database.DB.Where("user_id = ? AND platform = ?", userID, "jingtan").First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("未绑定鲸探账户")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	// TODO: 实际项目中需要调用鲸探API获取用户资产列表
	// 这里模拟返回一些资产

	// 模拟数据
	mockAssets := []struct {
		JingtanAssetID string
		Name           string
		ImageURL       string
	}{
		{"JT001", "鲸探限定藏品#1", "https://example.com/jingtan1.jpg"},
		{"JT002", "鲸探限定藏品#2", "https://example.com/jingtan2.jpg"},
	}

	var assets []models.JingtanAsset
	for _, mockAsset := range mockAssets {
		// 检查是否已存在
		var existingAsset models.JingtanAsset
		result := database.DB.Where("user_id = ? AND jingtan_asset_id = ?", userID, mockAsset.JingtanAssetID).First(&existingAsset)
		
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 不存在，创建新记录
			asset := models.JingtanAsset{
				UserID:         userID,
				JingtanAssetID: mockAsset.JingtanAssetID,
				Name:           mockAsset.Name,
				ImageURL:       mockAsset.ImageURL,
				Status:         "active",
			}
			if err := database.DB.Create(&asset).Error; err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		} else if result.Error == nil {
			// 已存在，更新信息
			existingAsset.Name = mockAsset.Name
			existingAsset.ImageURL = mockAsset.ImageURL
			if err := database.DB.Save(&existingAsset).Error; err != nil {
				return nil, err
			}
			assets = append(assets, existingAsset)
		} else {
			return nil, result.Error
		}
	}

	// 记录社区事件
	eventService := NewEventService()
	eventService.CreateEvent("jingtan_sync", userID, fmt.Sprintf("用户同步了 %d 个鲸探资产", len(assets)), userID, "user")

	return assets, nil
}

// GetUserJingtanAssets 获取用户的鲸探资产列表
func (s *JingtanService) GetUserJingtanAssets(userID uint64) ([]models.JingtanAsset, error) {
	var assets []models.JingtanAsset
	if err := database.DB.Where("user_id = ? AND status = ?", userID, "active").Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}
