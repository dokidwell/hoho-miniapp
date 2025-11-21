package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hoho-miniapp/backend/services"
)

type CreationHandler struct {
	creationService *services.CreationService
}

func NewCreationHandler(creationService *services.CreationService) *CreationHandler {
	return &CreationHandler{
		creationService: creationService,
	}
}

// SubmitCreation 提交创作
func (h *CreationHandler) SubmitCreation(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var req struct {
		Title        string `json:"title" binding:"required"`
		Description  string `json:"description"`
		MediaURL     string `json:"media_url" binding:"required"`
		ThumbnailURL string `json:"thumbnail_url"`
		TotalSupply  uint   `json:"total_supply" binding:"required,min=1"`
		Price        string `json:"price" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	creation, err := h.creationService.SubmitCreation(
		userID.(uint),
		req.Title,
		req.Description,
		req.MediaURL,
		req.ThumbnailURL,
		req.TotalSupply,
		req.Price,
	)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":  "创作提交成功，等待审核",
		"creation": creation,
	})
}

// GetMyCreations 获取我的创作列表
func (h *CreationHandler) GetMyCreations(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	creations, total, err := h.creationService.GetUserCreations(userID.(uint), status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"creations": creations,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetCreationDetail 获取创作详情
func (h *CreationHandler) GetCreationDetail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	creation, err := h.creationService.GetCreationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "创作不存在"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"creation": creation})
}
