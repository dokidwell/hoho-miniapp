package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hoho-miniapp/backend/services"
)

type OfferHandler struct {
	offerService *services.OfferService
}

func NewOfferHandler(offerService *services.OfferService) *OfferHandler {
	return &OfferHandler{
		offerService: offerService,
	}
}

// CreateOffer 创建出价
func (h *OfferHandler) CreateOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var req struct {
		ArtworkInstanceID uint   `json:"artwork_instance_id" binding:"required"`
		Price             string `json:"price" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	offer, err := h.offerService.CreateOffer(userID.(uint), req.ArtworkInstanceID, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "出价成功",
		"offer":   offer,
	})
}

// GetMyOffers 获取我的出价列表
func (h *OfferHandler) GetMyOffers(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	offers, total, err := h.offerService.GetUserOffers(userID.(uint), status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"offers":    offers,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// CancelOffer 取消出价
func (h *OfferHandler) CancelOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	offerID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := h.offerService.CancelOffer(uint(offerID), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "出价已取消"})
}

// AcceptOffer 接受出价
func (h *OfferHandler) AcceptOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	offerID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := h.offerService.AcceptOffer(uint(offerID), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "出价已接受"})
}
