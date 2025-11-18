package services

import (
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/models"
)

// EventService 定义社区事件服务接口
type EventService struct{}

// NewEventService 创建一个新的EventService实例
func NewEventService() *EventService {
	return &EventService{}
}

// CreateEvent 创建社区事件
func (s *EventService) CreateEvent(eventType string, userID uint64, description string, relatedID uint64, relatedType string) (*models.CommunityEvent, error) {
	event := models.CommunityEvent{
		EventType:   eventType,
		UserID:      userID,
		Description: description,
		RelatedID:   relatedID,
		RelatedType: relatedType,
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

// GetEvents 获取社区事件列表（分页）
func (s *EventService) GetEvents(page, pageSize int, eventType string) ([]models.CommunityEvent, int64, error) {
	var events []models.CommunityEvent
	var total int64

	query := database.DB.Model(&models.CommunityEvent{})

	// 如果指定了事件类型，则筛选
	if eventType != "" {
		query = query.Where("event_type = ?", eventType)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// GetEventByID 获取单个事件详情
func (s *EventService) GetEventByID(eventID uint64) (*models.CommunityEvent, error) {
	var event models.CommunityEvent
	if err := database.DB.First(&event, eventID).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
