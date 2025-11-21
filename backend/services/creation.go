package services

import (
	"errors"
	"fmt"
	"time"

	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
)

type CreationService struct{}

func NewCreationService() *CreationService {
	return &CreationService{}
}

// SubmitCreation 提交创作
func (s *CreationService) SubmitCreation(userID uint, title, description, mediaURL, thumbnailURL string, totalSupply uint, price string) (*models.Creation, error) {
	creation := &models.Creation{
		UserID:       userID,
		Title:        title,
		Description:  description,
		MediaURL:     mediaURL,
		ThumbnailURL: thumbnailURL,
		TotalSupply:  totalSupply,
		Price:        price,
		Status:       "pending",
	}
	
	// 获取默认分成比例
	var config models.SystemConfig
	if err := database.DB.Where("`key` = ?", "default_commission_rate").First(&config).Error; err == nil {
		creation.CommissionRate = config.Value
	}
	
	if err := database.DB.Create(creation).Error; err != nil {
		return nil, err
	}
	
	return creation, nil
}

// GetCreationByID 根据ID获取创作
func (s *CreationService) GetCreationByID(id uint) (*models.Creation, error) {
	var creation models.Creation
	if err := database.DB.Preload("User").First(&creation, id).Error; err != nil {
		return nil, err
	}
	return &creation, nil
}

// GetUserCreations 获取用户的创作列表
func (s *CreationService) GetUserCreations(userID uint, status string, page, pageSize int) ([]models.Creation, int64, error) {
	var creations []models.Creation
	var total int64
	
	query := database.DB.Where("user_id = ?", userID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	
	query.Model(&models.Creation{}).Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&creations).Error; err != nil {
		return nil, 0, err
	}
	
	return creations, total, nil
}

// GetPendingCreations 获取待审核的创作列表
func (s *CreationService) GetPendingCreations(page, pageSize int) ([]models.Creation, int64, error) {
	var creations []models.Creation
	var total int64
	
	query := database.DB.Where("status = ?", "pending").Preload("User")
	query.Model(&models.Creation{}).Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("created_at ASC").Offset(offset).Limit(pageSize).Find(&creations).Error; err != nil {
		return nil, 0, err
	}
	
	return creations, total, nil
}

// ApproveCreation 审核通过创作
func (s *CreationService) ApproveCreation(creationID, reviewerID uint) error {
	now := time.Now()
	
	// 更新创作状态
	if err := database.DB.Model(&models.Creation{}).Where("id = ?", creationID).Updates(map[string]interface{}{
		"status":      "approved",
		"reviewer_id": reviewerID,
		"reviewed_at": now,
	}).Error; err != nil {
		return err
	}
	
	// 获取创作信息
	var creation models.Creation
	if err := database.DB.Preload("User").First(&creation, creationID).Error; err != nil {
		return err
	}
	
	// 检查是否是首次创作通过
	var count int64
	database.DB.Model(&models.Creation{}).Where("user_id = ? AND status = ?", creation.UserID, "approved").Count(&count)
	
	if count == 1 {
		// 首次创作通过，发放奖励
		var task models.Task
		if err := database.DB.Where("condition_type = ? AND is_enabled = ?", "first_creation", true).First(&task).Error; err == nil {
			// 记录任务完成
			completion := &models.UserTaskCompletion{
				UserID:       creation.UserID,
				TaskID:       task.ID,
				CompletedAt:  now,
				RewardPoints: task.RewardPoints,
			}
			database.DB.Create(completion)
			
			// 发送通知
			notification := &models.Notification{
				UserID:  creation.UserID,
				Type:    "task",
				Title:   "任务完成",
				Content: fmt.Sprintf("恭喜！您的首次创作已通过审核，获得 %s 积分奖励", task.RewardPoints),
			}
			database.DB.Create(notification)
		}
	}
	
	// 发送审核通过通知
	notification := &models.Notification{
		UserID:    creation.UserID,
		Type:      "system",
		Title:     "创作审核通过",
		Content:   fmt.Sprintf("您的创作《%s》已通过审核", creation.Title),
		RelatedID: &creationID,
	}
	database.DB.Create(notification)
	
	return nil
}

// RejectCreation 审核拒绝创作
func (s *CreationService) RejectCreation(creationID, reviewerID uint, reason string) error {
	now := time.Now()
	
	// 更新创作状态
	if err := database.DB.Model(&models.Creation{}).Where("id = ?", creationID).Updates(map[string]interface{}{
		"status":        "rejected",
		"reviewer_id":   reviewerID,
		"reviewed_at":   now,
		"reject_reason": reason,
	}).Error; err != nil {
		return err
	}
	
	// 获取创作信息
	var creation models.Creation
	if err := database.DB.First(&creation, creationID).Error; err != nil {
		return err
	}
	
	// 发送审核拒绝通知
	notification := &models.Notification{
		UserID:    creation.UserID,
		Type:      "system",
		Title:     "创作审核未通过",
		Content:   fmt.Sprintf("您的创作《%s》未通过审核，原因：%s", creation.Title, reason),
		RelatedID: &creationID,
	}
	database.DB.Create(notification)
	
	return nil
}

// PublishCreation 发布创作
func (s *CreationService) PublishCreation(creationID uint, releaseDate *time.Time) error {
	// 检查创作状态
	var creation models.Creation
	if err := database.DB.First(&creation, creationID).Error; err != nil {
		return err
	}
	
	if creation.Status != "approved" {
		return errors.New("只有审核通过的创作才能发布")
	}
	
	now := time.Now()
	
	// 创建作品
	artwork := &models.Artwork{
		Title:        creation.Title,
		Description:  creation.Description,
		MediaType:    "image",
		MediaURL:     creation.MediaURL,
		ThumbnailURL: creation.ThumbnailURL,
		TotalSupply:  creation.TotalSupply,
		MintedCount:  0,
		Price:        creation.Price,
		Source:       "community",
		ReleaseDate:  releaseDate,
		Status:       "active",
	}
	
	if err := database.DB.Create(artwork).Error; err != nil {
		return err
	}
	
	// 更新创作状态
	if err := database.DB.Model(&models.Creation{}).Where("id = ?", creationID).Updates(map[string]interface{}{
		"status":       "published",
		"published_at": now,
	}).Error; err != nil {
		return err
	}
	
	// 发送发布通知
	notification := &models.Notification{
		UserID:    creation.UserID,
		Type:      "system",
		Title:     "作品已发布",
		Content:   fmt.Sprintf("您的创作《%s》已成功发布", creation.Title),
		RelatedID: &creationID,
	}
	database.DB.Create(notification)
	
	return nil
}
