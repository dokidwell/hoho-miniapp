package handlers

import (
	"hoho-miniapp/backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EventHandler 定义社区事件相关的HTTP处理函数
type EventHandler struct {
	EventService *services.EventService
}

// NewEventHandler 创建一个新的EventHandler实例
func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{EventService: eventService}
}

// ListEvents 处理获取社区事件列表请求
func (h *EventHandler) ListEvents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	eventType := c.DefaultQuery("event_type", "")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	events, total, err := h.EventService.GetEvents(page, pageSize, eventType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取事件列表失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": gin.H{
			"events": events,
			"total": total,
			"page": page,
			"page_size": pageSize,
		},
	})
}

// GetEventDetail 处理获取事件详情请求
func (h *EventHandler) GetEventDetail(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "事件ID格式错误"})
		return
	}

	event, err := h.EventService.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "事件不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": event,
	})
}
