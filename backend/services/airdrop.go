package services

import (
	"errors"
	"fmt"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// AirdropService 定义空投服务接口
type AirdropService struct{}

// NewAirdropService 创建一个新的AirdropService实例
func NewAirdropService() *AirdropService {
	return &AirdropService{}
}

// AirdropPoints 空投积分给用户
func (s *AirdropService) AirdropPoints(userID uint64, amount decimal.Decimal, reason string) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("空投金额必须大于0")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 更新用户积分
		if err := tx.Model(&models.UserPoint{}).Where("user_id = ?", userID).Updates(map[string]interface{}{
			"balance":      gorm.Expr("balance + ?", amount),
			"total_earned": gorm.Expr("total_earned + ?", amount),
		}).Error; err != nil {
			return err
		}

		// 2. 记录积分交易
		transaction := models.PointTransaction{
			UserID:          userID,
			Amount:          amount,
			TransactionType: "airdrop",
			Description:     reason,
			Status:          "completed",
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// 3. 记录社区事件
		event := models.CommunityEvent{
			EventType:   "airdrop_points",
			UserID:      userID,
			Description: fmt.Sprintf("用户获得空投积分 %.8f，原因：%s", amount, reason),
			RelatedID:   transaction.ID,
			RelatedType: "point_transaction",
		}
		if err := tx.Create(&event).Error; err != nil {
			return err
		}

		return nil
	})
}

// BatchAirdropPoints 批量空投积分
func (s *AirdropService) BatchAirdropPoints(userIDs []uint64, amount decimal.Decimal, reason string) (int, error) {
	if amount.LessThanOrEqual(decimal.Zero) {
		return 0, errors.New("空投金额必须大于0")
	}

	successCount := 0
	for _, userID := range userIDs {
		if err := s.AirdropPoints(userID, amount, reason); err != nil {
			// 记录错误但继续处理其他用户
			fmt.Printf("空投积分给用户 %d 失败: %v\n", userID, err)
			continue
		}
		successCount++
	}

	return successCount, nil
}

// AirdropAsset 空投藏品给用户（基于已有的Asset）
func (s *AirdropService) AirdropAsset(assetID uint64, targetUserID uint64, count int) ([]models.AssetInstance, error) {
	// 复用AssetService的MintAndAirdrop方法
	assetService := NewAssetService()
	return assetService.MintAndAirdrop(assetID, targetUserID, count)
}

// GetAirdropHistory 获取空投历史记录
func (s *AirdropService) GetAirdropHistory(page, pageSize int) ([]models.PointTransaction, int64, error) {
	var transactions []models.PointTransaction
	var total int64

	query := database.DB.Where("transaction_type = ?", "airdrop")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}
