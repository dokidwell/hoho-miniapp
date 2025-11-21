package services

import (
	"errors"
	"fmt"
	"time"

	"hoho-api/database"
	"hoho-api/models"
	
	"github.com/shopspring/decimal"
)

// CreateOffer 创建出价
func CreateOffer(buyerID, artworkInstanceID uint, price string) (*models.Offer, error) {
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
	
	// 检查积分余额
	priceDecimal, _ := decimal.NewFromString(price)
	balance, err := GetAvailablePoints(buyerID)
	if err != nil {
		return nil, err
	}
	
	if balance.LessThan(priceDecimal) {
		return nil, errors.New("积分余额不足")
	}
	
	// 冻结积分
	if err := FreezePoints(buyerID, priceDecimal, "offer", "出价冻结", nil); err != nil {
		return nil, err
	}
	
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
		// 创建失败，解冻积分
		UnfreezePoints(buyerID, priceDecimal, "offer_cancel", "出价取消，解冻积分", nil)
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
func GetUserOffers(userID uint, status string, page, pageSize int) ([]models.Offer, int64, error) {
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
func GetReceivedOffers(ownerID uint, status string, page, pageSize int) ([]models.Offer, int64, error) {
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
func AcceptOffer(offerID, sellerID uint) error {
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
		// 过期，解冻积分
		price, _ := decimal.NewFromString(offer.Price)
		UnfreezePoints(offer.BuyerID, price, "offer_expired", "出价过期，解冻积分", &offerID)
		
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
	
	// 解冻买家积分
	price, _ := decimal.NewFromString(offer.Price)
	if err := UnfreezePoints(offer.BuyerID, price, "offer_accept", "出价被接受，解冻积分", &offerID); err != nil {
		tx.Rollback()
		return err
	}
	
	// 扣除买家积分
	if err := DeductPoints(offer.BuyerID, price, "trade", fmt.Sprintf("购买作品：%s", offer.ArtworkInstance.Artwork.Title), &offerID); err != nil {
		tx.Rollback()
		return err
	}
	
	// 计算手续费
	var feeConfig models.SystemConfig
	feeRate := decimal.NewFromFloat(2.0) // 默认2%
	if err := tx.Where("`key` = ?", "trade_fee_rate").First(&feeConfig).Error; err == nil {
		feeRate, _ = decimal.NewFromString(feeConfig.Value)
	}
	
	fee := price.Mul(feeRate).Div(decimal.NewFromInt(100))
	sellerReceive := price.Sub(fee)
	
	// 增加卖家积分
	if err := AddPoints(sellerID, sellerReceive, "trade", fmt.Sprintf("出售作品：%s", offer.ArtworkInstance.Artwork.Title), &offerID); err != nil {
		tx.Rollback()
		return err
	}
	
	// 记录平台手续费收入
	RecordPlatformIncome("fee", fee.String(), fmt.Sprintf("交易手续费：作品 #%d", offer.ArtworkInstanceID), &offerID)
	
	// 转移作品所有权
	if err := tx.Model(&models.ArtworkInstance{}).Where("id = ?", offer.ArtworkInstanceID).
		Update("owner_id", offer.BuyerID).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 创建交易记录
	trade := &models.Trade{
		SellerID:          sellerID,
		BuyerID:           offer.BuyerID,
		ArtworkInstanceID: offer.ArtworkInstanceID,
		Price:             offer.Price,
		Fee:               fee.String(),
		Type:              "offer",
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
func RejectOffer(offerID, sellerID uint) error {
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
	
	// 解冻买家积分
	price, _ := decimal.NewFromString(offer.Price)
	if err := UnfreezePoints(offer.BuyerID, price, "offer_reject", "出价被拒绝，解冻积分", &offerID); err != nil {
		return err
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
		Content:   "您的出价已被拒绝，积分已退回",
		RelatedID: &offerID,
	}
	database.DB.Create(notification)
	
	return nil
}

// CancelOffer 取消出价
func CancelOffer(offerID, buyerID uint) error {
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
	
	// 解冻积分
	price, _ := decimal.NewFromString(offer.Price)
	if err := UnfreezePoints(buyerID, price, "offer_cancel", "取消出价，解冻积分", &offerID); err != nil {
		return err
	}
	
	// 更新出价状态
	if err := database.DB.Model(&offer).Update("status", "cancelled").Error; err != nil {
		return err
	}
	
	return nil
}

// ExpireOffers 过期出价处理（定时任务）
func ExpireOffers() error {
	var offers []models.Offer
	if err := database.DB.Where("status = ? AND expires_at < ?", "pending", time.Now()).Find(&offers).Error; err != nil {
		return err
	}
	
	for _, offer := range offers {
		// 解冻积分
		price, _ := decimal.NewFromString(offer.Price)
		UnfreezePoints(offer.BuyerID, price, "offer_expired", "出价过期，解冻积分", &offer.ID)
		
		// 更新状态
		database.DB.Model(&offer).Update("status", "expired")
		
		// 发送通知
		notification := &models.Notification{
			UserID:    offer.BuyerID,
			Type:      "offer",
			Title:     "出价已过期",
			Content:   "您的出价已过期，积分已退回",
			RelatedID: &offer.ID,
		}
		database.DB.Create(notification)
	}
	
	return nil
}
