package services

import (
	"context"
	"errors"
	"fmt"
	"hoho-miniapp/backend/config"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// TradeService 定义交易服务接口
type TradeService struct{}

// NewTradeService 创建一个新的TradeService实例
func NewTradeService() *TradeService {
	return &TradeService{}
}

// CreateListing 创建挂售单
func (s *TradeService) CreateListing(sellerID uint64, assetInstanceID uint64, price decimal.Decimal) (*models.Listing, error) {
	// 1. 检查藏品实例是否存在且属于卖家
	var instance models.AssetInstance
	if err := database.DB.First(&instance, assetInstanceID).Error; err != nil {
		return nil, errors.New("藏品实例不存在")
	}

	if instance.OwnerID != sellerID {
		return nil, errors.New("你不是该藏品的拥有者")
	}

	if instance.Status != "in_wallet" {
		return nil, errors.New("该藏品不可交易")
	}

	// 2. 创建Listing
	listing := models.Listing{
		AssetInstanceID: assetInstanceID,
		SellerID:        sellerID,
		Price:           price,
		Status:          "active",
	}

	if err := database.DB.Create(&listing).Error; err != nil {
		return nil, err
	}

	// 3. 更新AssetInstance状态为on_sale
	if err := database.DB.Model(&instance).Update("status", "on_sale").Error; err != nil {
		// 回滚Listing创建
		database.DB.Delete(&listing)
		return nil, err
	}

	return &listing, nil
}

// CancelListing 取消挂售单
func (s *TradeService) CancelListing(listingID uint64, userID uint64) error {
	var listing models.Listing
	if err := database.DB.First(&listing, listingID).Error; err != nil {
		return errors.New("挂售单不存在")
	}

	if listing.SellerID != userID {
		return errors.New("你没有权限取消此挂售单")
	}

	if listing.Status != "active" {
		return errors.New("该挂售单已处理")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 更新Listing状态
		if err := tx.Model(&listing).Update("status", "canceled").Error; err != nil {
			return err
		}

		// 2. 更新AssetInstance状态回到in_wallet
		if err := tx.Model(&models.AssetInstance{}).Where("id = ?", listing.AssetInstanceID).Update("status", "in_wallet").Error; err != nil {
			return err
		}

		return nil
	})
}

// ExecuteTrade 执行交易（核心逻辑，包含并发锁）
func (s *TradeService) ExecuteTrade(listingID uint64, buyerID uint64) (*models.Trade, error) {
	var listing models.Listing
	var instance models.AssetInstance
	var asset models.Asset
	var buyerPoints models.UserPoint
	var sellerPoints models.UserPoint

	// 1. 获取Listing
	if err := database.DB.First(&listing, listingID).Error; err != nil {
		return nil, errors.New("挂售单不存在")
	}

	if listing.Status != "active" {
		return nil, errors.New("该挂售单已处理")
	}

	// 2. 获取藏品实例
	if err := database.DB.First(&instance, listing.AssetInstanceID).Error; err != nil {
		return nil, errors.New("藏品实例不存在")
	}

	// 3. 获取藏品信息（用于获取创作者信息）
	if err := database.DB.First(&asset, instance.AssetID).Error; err != nil {
		return nil, errors.New("藏品不存在")
	}

	// 4. 获取买家和卖家的积分信息
	if err := database.DB.Where("user_id = ?", buyerID).First(&buyerPoints).Error; err != nil {
		return nil, errors.New("买家积分信息不存在")
	}

	if err := database.DB.Where("user_id = ?", listing.SellerID).First(&sellerPoints).Error; err != nil {
		return nil, errors.New("卖家积分信息不存在")
	}

	// 5. **使用Redis分布式锁**确保并发安全
	lockKey := fmt.Sprintf("trade:lock:%d", listing.AssetInstanceID)
	lockValue := fmt.Sprintf("%d:%d", buyerID, time.Now().UnixNano())

	// 尝试获取锁（最多重试3次）
	var lockAcquired bool
	for i := 0; i < 3; i++ {
		ok, err := database.RDB.SetNX(context.Background(), lockKey, lockValue, 30*time.Second).Result()
		if err != nil {
			return nil, fmt.Errorf("获取分布式锁失败: %w", err)
		}
		if ok {
			lockAcquired = true
			break
		}
		time.Sleep(100 * time.Millisecond) // 等待后重试
	}

	if !lockAcquired {
		return nil, errors.New("该藏品正在交易中，请稍后重试")
	}

	// 确保锁释放
	defer func() {
		// 验证锁值，防止误删其他请求的锁
		val, _ := database.RDB.Get(context.Background(), lockKey).Result()
		if val == lockValue {
			database.RDB.Del(context.Background(), lockKey)
		}
	}()

	// 6. **在事务中执行交易**
	trade := &models.Trade{}
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 6.1. 再次检查Listing和Instance状态（防止竞态条件）
		var freshListing models.Listing
		if err := tx.First(&freshListing, listingID).Error; err != nil {
			return errors.New("挂售单不存在")
		}
		if freshListing.Status != "active" {
			return errors.New("该挂售单已被其他用户购买")
		}

		var freshInstance models.AssetInstance
		if err := tx.First(&freshInstance, listing.AssetInstanceID).Error; err != nil {
			return errors.New("藏品实例不存在")
		}
		if freshInstance.Status != "on_sale" {
			return errors.New("该藏品状态已变更，无法购买")
		}

		// 6.2. 检查买家积分是否足够
		freshBuyerPoints := models.UserPoint{}
		if err := tx.Where("user_id = ?", buyerID).First(&freshBuyerPoints).Error; err != nil {
			return errors.New("买家积分信息不存在")
		}

		if freshBuyerPoints.Balance.LessThan(listing.Price) {
			return errors.New("买家积分不足")
		}

		// 6.3. 计算手续费和版税（使用银行家舍入法，精确到8位小数）
		platformFee := listing.Price.Mul(config.AppConfig.PlatformFeeRate).RoundBank(config.AppConfig.DecimalPrecision)
		creatorRoyalty := listing.Price.Mul(config.AppConfig.CreatorRoyaltyRate).RoundBank(config.AppConfig.DecimalPrecision)
		sellerReceived := listing.Price.Sub(platformFee).Sub(creatorRoyalty)

		// 验证积分守恒（防止舍入误差导致积分不守恒）
		totalDistributed := platformFee.Add(creatorRoyalty).Add(sellerReceived)
		if !totalDistributed.Equal(listing.Price) {
			// 如果有舍入误差，调整卖家收入（误差通常在0.00000001以内）
			diff := listing.Price.Sub(totalDistributed)
			sellerReceived = sellerReceived.Add(diff)
		}

		// 6.4. 创建Trade记录
		*trade = models.Trade{
			ListingID:       listingID,
			AssetInstanceID: listing.AssetInstanceID,
			BuyerID:         buyerID,
			SellerID:        listing.SellerID,
			Price:           listing.Price,
			PlatformFee:     platformFee,
			CreatorRoyalty:  creatorRoyalty,
			SellerReceived:  sellerReceived,
			Status:          "pending",
		}

		if err := tx.Create(trade).Error; err != nil {
			return err
		}

		// 6.5. 冻结买家积分
		if err := tx.Model(&models.UserPoint{}).Where("user_id = ?", buyerID).Update("frozen", freshBuyerPoints.Frozen.Add(listing.Price)).Error; err != nil {
			return err
		}

		// 6.6. 更新Listing状态为sold
		if err := tx.Model(&freshListing).Update("status", "sold").Error; err != nil {
			return err
		}

		// 6.7. 更新AssetInstance状态为pending_trade
		if err := tx.Model(&freshInstance).Update("status", "pending_trade").Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 7. **异步处理积分转移**（这里简化为同步，实际应该用消息队列）
	if err := s.CompleteTradePayment(trade.ID); err != nil {
		// 记录错误但不中断，因为交易记录已创建
		fmt.Printf("警告：交易%d的积分转移失败: %v\n", trade.ID, err)
	}

	return trade, nil
}

// CompleteTradePayment 完成交易的积分转移
func (s *TradeService) CompleteTradePayment(tradeID uint64) error {
	var trade models.Trade
	if err := database.DB.First(&trade, tradeID).Error; err != nil {
		return errors.New("交易不存在")
	}

	if trade.Status != "pending" {
		return errors.New("交易已处理")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 买家：扣减积分（从冻结改为已扣减）
		if err := tx.Model(&models.UserPoint{}).Where("user_id = ?", trade.BuyerID).Updates(map[string]interface{}{
			"balance":     gorm.Expr("balance - ?", trade.Price),
			"frozen":      gorm.Expr("frozen - ?", trade.Price),
			"total_spent": gorm.Expr("total_spent + ?", trade.Price),
		}).Error; err != nil {
			return err
		}

		// 2. 卖家：增加积分
		if err := tx.Model(&models.UserPoint{}).Where("user_id = ?", trade.SellerID).Updates(map[string]interface{}{
			"balance":      gorm.Expr("balance + ?", trade.SellerReceived),
			"total_earned": gorm.Expr("total_earned + ?", trade.SellerReceived),
		}).Error; err != nil {
			return err
		}

		// 3. 创作者：增加版税
		if err := tx.Model(&models.UserPoint{}).Where("user_id = ?", gorm.Expr("(SELECT creator_id FROM assets WHERE id = (SELECT asset_id FROM asset_instances WHERE id = ?))", trade.AssetInstanceID)).Updates(map[string]interface{}{
			"balance":      gorm.Expr("balance + ?", trade.CreatorRoyalty),
			"total_earned": gorm.Expr("total_earned + ?", trade.CreatorRoyalty),
		}).Error; err != nil {
			return err
		}

		// 4. 更新AssetInstance的所有者
		if err := tx.Model(&models.AssetInstance{}).Where("id = ?", trade.AssetInstanceID).Updates(map[string]interface{}{
			"owner_id": trade.BuyerID,
			"status":   "in_wallet",
		}).Error; err != nil {
			return err
		}

		// 5. 更新Trade状态为completed
		if err := tx.Model(&trade).Update("status", "completed").Error; err != nil {
			return err
		}

		// 6. 记录社区事件
		event := models.CommunityEvent{
			EventType:   "trade",
			UserID:      trade.BuyerID,
			Description: fmt.Sprintf("用户 uid%d 以 %s 积分购买了藏品", trade.BuyerID, trade.Price.String()),
			RelatedID:   trade.ID,
			RelatedType: "trade",
		}
		if err := tx.Create(&event).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetListings 获取挂售单列表
func (s *TradeService) GetListings(page, pageSize int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var total int64

	query := database.DB.Where("status = ?", "active")
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&listings).Error; err != nil {
		return nil, 0, err
	}

	return listings, total, nil
}

// ListListings 获取挂售列表（支持状态筛选）
func (s *TradeService) ListListings(status string, page, pageSize int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var total int64

	query := database.DB.Model(&models.Listing{})

	// 如果指定了状态，则筛选
	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Preload("AssetInstance").Preload("AssetInstance.Asset").Limit(pageSize).Offset(offset).Order("id desc").Find(&listings).Error; err != nil {
		return nil, 0, err
	}

	return listings, total, nil
}

// GetListingDetail 获取挂售详情
func (s *TradeService) GetListingDetail(listingID uint64) (*models.Listing, error) {
	var listing models.Listing
	if err := database.DB.Preload("AssetInstance").Preload("AssetInstance.Asset").First(&listing, listingID).Error; err != nil {
		return nil, errors.New("挂售单不存在")
	}
	return &listing, nil
}

// GetTradeHistory 获取交易历史
func (s *TradeService) GetTradeHistory(userID uint64, tradeType string, page, pageSize int) ([]models.Trade, int64, error) {
	var trades []models.Trade
	var total int64

	query := database.DB.Model(&models.Trade{})

	// 根据类型筛选
	switch tradeType {
	case "buy":
		query = query.Where("buyer_id = ?", userID)
	case "sell":
		query = query.Where("seller_id = ?", userID)
	default:
		// all - 买入或卖出的都显示
		query = query.Where("buyer_id = ? OR seller_id = ?", userID, userID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Preload("AssetInstance").Preload("AssetInstance.Asset").Limit(pageSize).Offset(offset).Order("id desc").Find(&trades).Error; err != nil {
		return nil, 0, err
	}

	return trades, total, nil
}

// GetTradeDetail 获取交易详情
func (s *TradeService) GetTradeDetail(tradeID uint64, userID uint64) (*models.Trade, error) {
	var trade models.Trade
	if err := database.DB.Preload("AssetInstance").Preload("AssetInstance.Asset").First(&trade, tradeID).Error; err != nil {
		return nil, errors.New("交易不存在")
	}

	// 验证用户权限（只能查看自己参与的交易）
	if trade.BuyerID != userID && trade.SellerID != userID {
		return nil, errors.New("你没有权限查看此交易")
	}

	return &trade, nil
}

// GetMyListings 获取我的挂售列表
func (s *TradeService) GetMyListings(userID uint64, status string, page, pageSize int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var total int64

	query := database.DB.Model(&models.Listing{}).Where("seller_id = ?", userID)

	// 如果指定了状态，则筛选
	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Preload("AssetInstance").Preload("AssetInstance.Asset").Limit(pageSize).Offset(offset).Order("id desc").Find(&listings).Error; err != nil {
		return nil, 0, err
	}

	return listings, total, nil
}
