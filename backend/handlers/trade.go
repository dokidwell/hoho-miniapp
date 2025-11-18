package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"hoho-miniapp/backend/services"
)

// TradeHandler 定义交易处理器
type TradeHandler struct {
	tradeService *services.TradeService
}

// NewTradeHandler 创建一个新的TradeHandler实例
func NewTradeHandler(tradeService *services.TradeService) *TradeHandler {
	return &TradeHandler{
		tradeService: tradeService,
	}
}

// CreateListing 创建挂售单
// POST /api/v1/listings
func (h *TradeHandler) CreateListing(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var req struct {
		AssetInstanceID uint64 `json:"asset_instance_id" binding:"required"`
		Price           string `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 解析价格
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "价格格式错误"})
		return
	}

	// 价格必须大于0
	if price.LessThanOrEqual(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "价格必须大于0"})
		return
	}

	// 创建挂售单
	listing, err := h.tradeService.CreateListing(userID.(uint64), req.AssetInstanceID, price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "挂售成功",
		"data":    listing,
	})
}

// ListListings 获取挂售列表
// GET /api/v1/listings
func (h *TradeHandler) ListListings(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.DefaultQuery("status", "active") // active, sold, cancelled

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 获取挂售列表
	listings, total, err := h.tradeService.ListListings(status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": listings,
		"pagination": gin.H{
			"page":       page,
			"page_size":  pageSize,
			"total":      total,
			"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// GetListingDetail 获取挂售详情
// GET /api/v1/listings/:id
func (h *TradeHandler) GetListingDetail(c *gin.Context) {
	listingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的挂售ID"})
		return
	}

	listing, err := h.tradeService.GetListingDetail(listingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": listing,
	})
}

// CancelListing 取消挂售
// DELETE /api/v1/listings/:id
func (h *TradeHandler) CancelListing(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	listingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的挂售ID"})
		return
	}

	if err := h.tradeService.CancelListing(listingID, userID.(uint64)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "取消挂售成功",
	})
}

// ExecuteTrade 执行交易
// POST /api/v1/trades/execute
func (h *TradeHandler) ExecuteTrade(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var req struct {
		ListingID uint64 `json:"listing_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 执行交易
	trade, err := h.tradeService.ExecuteTrade(req.ListingID, userID.(uint64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "交易成功",
		"data":    trade,
	})
}

// GetTradeHistory 获取交易历史
// GET /api/v1/trades/history
func (h *TradeHandler) GetTradeHistory(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	tradeType := c.DefaultQuery("type", "") // buy, sell, all

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 获取交易历史
	trades, total, err := h.tradeService.GetTradeHistory(userID.(uint64), tradeType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trades,
		"pagination": gin.H{
			"page":       page,
			"page_size":  pageSize,
			"total":      total,
			"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// GetTradeDetail 获取交易详情
// GET /api/v1/trades/:id
func (h *TradeHandler) GetTradeDetail(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	tradeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的交易ID"})
		return
	}

	trade, err := h.tradeService.GetTradeDetail(tradeID, userID.(uint64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trade,
	})
}

// GetMyListings 获取我的挂售列表
// GET /api/v1/my/listings
func (h *TradeHandler) GetMyListings(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.DefaultQuery("status", "") // active, sold, cancelled, all

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 获取我的挂售列表
	listings, total, err := h.tradeService.GetMyListings(userID.(uint64), status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": listings,
		"pagination": gin.H{
			"page":       page,
			"page_size":  pageSize,
			"total":      total,
			"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}
