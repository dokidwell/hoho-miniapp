package services

import (
	"errors"
	"fmt"
	"time"

	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
	
	"github.com/shopspring/decimal"
)

type OfferService struct{}

func NewOfferService() *OfferService {
	return &OfferService{}
}

// CreateOffer 创建出价
func (s *OfferService) CreateOffer(buyerID, artworkInstanceID uint, price string) (*models.Offer, error) {
	// 检查作品实例是否存在
	var instance models.ArtworkInstance
	if err := database.DB.Preload("Artwork").Preload("Owner").First(&instance, artworkInstanceID).Error; err != nil {
		return nil, errors.New("作品不存在")
	}
	
	// 检查是否是自己的作品
	if instance.OwnerID == buyerID {
		return nil, errors.New("不能对自己的作品出价")
	}
	
	// 检查作品状态
	if instance.Status == "listed" {
		return nil, errors.New("该作品正在挂售中，请直接购买")
	}
	
	// TODO: 实现积分检查和冻结逻辑
	priceDecimal, _ := decimal.NewFromString(price)
	_ = priceDecimal
	
	// 获取出价有效期配置
	var config models.SystemConfig
	expireDays := 7 // 默认7天
	if err := database.DB.Where("`key` = ?", "offer_expire_days").First(&config).Error; err == nil {
		fmt.Sscanf(config.Value, "%d", &expireDays)
	}
	
	expiresAt := time.Now().Add(time.Duration(expireDays) * 24 * time.Hour)
	
	// 创建出价
	offer := &models.Offer{
		BuyerID:           buyerID,
		ArtworkInstanceID: artworkInstanceID,
		Price:             price,
		Status:            "pending",
		ExpiresAt:         &expiresAt,
	}
	
	if err := database.DB.Create(offer).Error; err != nil {
		return nil, err
	}
	
	// 发送通知给卖家
	notification := &models.Notification{
		UserID:    instance.OwnerID,
		Type:      "offer",
		Title:     "收到新出价",
		Content:   fmt.Sprintf("您的作品《%s》收到了 %s 积分的出价", instance.Artwork.Title, price),
		RelatedID: &offer.ID,
	}
	database.DB.Create(notification)
	
	return offer, nil
}

// GetUserOffers 获取用户的出价列表
func (s *OfferService) GetUserOffers(userID uint, status string, page, pageSize int) ([]models.Offer, int64, error) {
	var offers []models.Offer
	var total int64
	
	query := database.DB.Where("buyer_id = ?", userID).Preload("ArtworkInstance.Artwork")
	
	if status != "" {
		query = query.Where("status = ?", status)
	}
	
	query.Model(&models.Offer{}).Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&offers).Error; err != nil {
		return nil, 0, err
	}
	
	return offers, total, nil
}

// GetReceivedOffers 获取收到的出价列表
func (s *OfferService) GetReceivedOffers(ownerID uint, status string, page, pageSize int) ([]models.Offer, int64, error) {
	var offers []models.Offer
	var total int64
	
	query := database.DB.Joins("JOIN artwork_instances ON offers.artwork_instance_id = artwork_instances.id").
		Where("artwork_instances.owner_id = ?", ownerID).
		Preload("ArtworkInstance.Artwork").
		Preload("Buyer")
	
	if status != "" {
		query = query.Where("offers.status = ?", status)
	}
	
	query.Model(&models.Offer{}).Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("offers.created_at DESC").Offset(offset).Limit(pageSize).Find(&offers).Error; err != nil {
		return nil, 0, err
	}
	
	return offers, total, nil
}

// AcceptOffer 接受出价
func (s *OfferService) AcceptOffer(offerID, sellerID uint) error {
	// 获取出价信息
	var offer models.Offer
	if err := database.DB.Preload("ArtworkInstance.Artwork").Preload("Buyer").First(&offer, offerID).Error; err != nil {
		return err
	}
	
	// 检查出价状态
	if offer.Status != "pending" {
		return errors.New("出价已处理")
	}
	
	// 检查是否过期
	if offer.ExpiresAt != nil && time.Now().After(*offer.ExpiresAt) {
		database.DB.Model(&offer).Update("status", "expired")
		return errors.New("出价已过期")
	}
	
	// 检查是否是作品所有者
	if offer.ArtworkInstance.OwnerID != sellerID {
		return errors.New("无权接受该出价")
	}
	
	// 开始交易
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	// 计算手续费
	price, _ := decimal.NewFromString(offer.Price)
	var feeConfig models.SystemConfig
	feeRate := decimal.NewFromFloat(2.0) // 默认2%
	if err := tx.Where("`key` = ?", "trade_fee_rate").First(&feeConfig).Error; err == nil {
		feeRate, _ = decimal.NewFromString(feeConfig.Value)
	}
	
	fee := price.Mul(feeRate).Div(decimal.NewFromInt(100))
	
	// 转移作品所有权
	if err := tx.Model(&models.ArtworkInstance{}).Where("id = ?", offer.ArtworkInstanceID).
		Update("owner_id", offer.BuyerID).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 创建交易记录
	sellerReceived := price.Sub(fee)
	trade := &models.Trade{
		AssetInstanceID: uint64(offer.ArtworkInstanceID),
		SellerID:        uint64(sellerID),
		BuyerID:         uint64(offer.BuyerID),
		Price:           price,
		PlatformFee:     fee,
		CreatorRoyalty:  decimal.Zero,
		SellerReceived:  sellerReceived,
		Status:          "completed",
	}
	if err := tx.Create(trade).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 更新出价状态
	now := time.Now()
	if err := tx.Model(&offer).Updates(map[string]interface{}{
		"status":       "accepted",
		"responded_at": now,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	tx.Commit()
	
	// 发送通知
	buyerNotification := &models.Notification{
		UserID:    offer.BuyerID,
		Type:      "offer",
		Title:     "出价已接受",
		Content:   fmt.Sprintf("您对作品《%s》的出价已被接受", offer.ArtworkInstance.Artwork.Title),
		RelatedID: &offerID,
	}
	database.DB.Create(buyerNotification)
	
	return nil
}

// RejectOffer 拒绝出价
func (s *OfferService) RejectOffer(offerID, sellerID uint) error {
	// 获取出价信息
	var offer models.Offer
	if err := database.DB.Preload("ArtworkInstance").First(&offer, offerID).Error; err != nil {
		return err
	}
	
	// 检查出价状态
	if offer.Status != "pending" {
		return errors.New("出价已处理")
	}
	
	// 检查是否是作品所有者
	if offer.ArtworkInstance.OwnerID != sellerID {
		return errors.New("无权拒绝该出价")
	}
	
	// 更新出价状态
	now := time.Now()
	if err := database.DB.Model(&offer).Updates(map[string]interface{}{
		"status":       "rejected",
		"responded_at": now,
	}).Error; err != nil {
		return err
	}
	
	// 发送通知
	notification := &models.Notification{
		UserID:    offer.BuyerID,
		Type:      "offer",
		Title:     "出价已拒绝",
		Content:   "您的出价已被拒绝",
		RelatedID: &offerID,
	}
	database.DB.Create(notification)
	
	return nil
}

// CancelOffer 取消出价
func (s *OfferService) CancelOffer(offerID, buyerID uint) error {
	// 获取出价信息
	var offer models.Offer
	if err := database.DB.First(&offer, offerID).Error; err != nil {
		return err
	}
	
	// 检查是否是出价者
	if offer.BuyerID != buyerID {
		return errors.New("无权取消该出价")
	}
	
	// 检查出价状态
	if offer.Status != "pending" {
		return errors.New("出价已处理，无法取消")
	}
	
	// 更新出价状态
	if err := database.DB.Model(&offer).Update("status", "cancelled").Error; err != nil {
		return err
	}
	
	return nil
}

// ExpireOffers 过期出价处理（定时任务）
func (s *OfferService) ExpireOffers() error {
	var offers []models.Offer
	if err := database.DB.Where("status = ? AND expires_at < ?", "pending", time.Now()).Find(&offers).Error; err != nil {
		return err
	}
	
	for _, offer := range offers {
		// 更新状态
		database.DB.Model(&offer).Update("status", "expired")
		
		// 发送通知
		notification := &models.Notification{
			UserID:    offer.BuyerID,
			Type:      "offer",
			Title:     "出价已过期",
			Content:   "您的出价已过期",
			RelatedID: &offer.ID,
		}
		database.DB.Create(notification)
	}
	
	return nil
}
