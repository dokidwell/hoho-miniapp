package services

import (
	"errors"
	"fmt"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
	"hoho-miniapp/backend/utils"
	"time"

	"gorm.io/gorm"
)

// AssetService 定义藏品服务接口
type AssetService struct{}

// NewAssetService 创建一个新的AssetService实例
func NewAssetService() *AssetService {
	return &AssetService{}
}

// CreateCollection 创建藏品集合
func (s *AssetService) CreateCollection(name, description, coverImage string) (*models.Collection, error) {
	collection := models.Collection{
		Name:        name,
		Description: description,
		CoverImage:  coverImage,
		Status:      "active",
	}
	if err := database.DB.Create(&collection).Error; err != nil {
		return nil, err
	}
	return &collection, nil
}

// SubmitMintRequest 用户提交铸造请求
func (s *AssetService) SubmitMintRequest(creatorID uint64, collectionID uint64, name, description, mediaURL, mediaType string, totalSupply int) (*models.Asset, error) {
	// 1. 检查集合是否存在
	var collection models.Collection
	if err := database.DB.First(&collection, collectionID).Error; err != nil {
		return nil, errors.New("藏品集合不存在")
	}

	// 2. 创建Asset（藏品SKU），状态为待审核
	asset := models.Asset{
		CollectionID: collectionID,
		Name:         name,
		Description:  description,
		MediaURL:     mediaURL,
		MediaType:    mediaType,
		TotalSupply:  totalSupply,
		CreatorID:    creatorID,
		Status:       "pending_review",
	}

	if err := database.DB.Create(&asset).Error; err != nil {
		return nil, err
	}

	return &asset, nil
}

// ReviewMintRequest 审核铸造请求
func (s *AssetService) ReviewMintRequest(assetID uint64, status string) (*models.Asset, error) {
	var asset models.Asset
	if err := database.DB.First(&asset, assetID).Error; err != nil {
		return nil, errors.New("铸造请求不存在")
	}

	if asset.Status != "pending_review" {
		return nil, errors.New("该请求已处理")
	}

	if status == "approved" {
		// 审核通过，将状态改为active
		if err := database.DB.Model(&asset).Updates(models.Asset{Status: "active"}).Error; err != nil {
			return nil, err
		}
		// 自动铸造所有实例并空投给创作者
		if _, err := s.MintAndAirdrop(asset.ID, asset.CreatorID, asset.TotalSupply); err != nil {
			// 铸造失败，需要回滚状态或记录错误
			database.DB.Model(&asset).Updates(models.Asset{Status: "rejected"}) // 简单回滚
			return nil, fmt.Errorf("铸造失败: %w", err)
		}
	} else if status == "rejected" {
		// 审核拒绝
		if err := database.DB.Model(&asset).Updates(models.Asset{Status: "rejected"}).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("无效的审核状态")
	}

	return &asset, nil
}

// MintAndAirdrop 铸造藏品实例并空投给指定用户
func (s *AssetService) MintAndAirdrop(assetID uint64, targetUserID uint64, count int) ([]models.AssetInstance, error) {
	var asset models.Asset
	if err := database.DB.First(&asset, assetID).Error; err != nil {
		return nil, errors.New("藏品不存在")
	}

	if asset.Status != "active" {
		return nil, errors.New("藏品未激活或已禁用")
	}

	if asset.MintedCount+count > asset.TotalSupply {
		return nil, errors.New("铸造数量超过总发行量")
	}

	var instances []models.AssetInstance
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < count; i++ {
			// 实例编号从已铸造数量开始递增
			instanceNo := asset.MintedCount + i + 1

			instance := models.AssetInstance{
				AssetID:    assetID,
				InstanceNo: instanceNo,
				OwnerID:    targetUserID,
				TokenID:    utils.GenerateTokenID(assetID, uint64(instanceNo)), // 生成唯一TokenID
				Status:     "in_wallet",
			}
			if err := tx.Create(&instance).Error; err != nil {
				return err
			}
			instances = append(instances, instance)
		}

		// 更新Asset的已铸造数量
		if err := tx.Model(&asset).Update("minted_count", asset.MintedCount+count).Error; err != nil {
			return err
		}

		// TODO: 记录社区事件（空投事件）

		return nil
	})

	if err != nil {
		return nil, err
	}

	return instances, nil
}

// GetPendingReviewAssets 获取待审核的铸造请求
func (s *AssetService) GetPendingReviewAssets() ([]models.Asset, error) {
	var assets []models.Asset
	err := database.DB.Where("status = ?", "pending_review").Find(&assets).Error
	return assets, err
}

// GetAssetByID 获取藏品SKU详情
func (s *AssetService) GetAssetByID(assetID uint64) (*models.Asset, error) {
	var asset models.Asset
	err := database.DB.First(&asset, assetID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("藏品不存在")
	}
	return &asset, err
}
