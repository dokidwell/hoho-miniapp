package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hoho-miniapp/backend/services"
)

type PlatformAccountHandler struct {
	platformAccountService *services.PlatformAccountService
}

func NewPlatformAccountHandler(platformAccountService *services.PlatformAccountService) *PlatformAccountHandler {
	return &PlatformAccountHandler{
		platformAccountService: platformAccountService,
	}
}

// GetAccountInfo 获取平台账户信息（阳光账户）
func (h *PlatformAccountHandler) GetAccountInfo(c *gin.Context) {
	account, err := h.platformAccountService.GetPlatformAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"account": account})
}

// GetTransactions 获取平台交易记录
func (h *PlatformAccountHandler) GetTransactions(c *gin.Context) {
	transactionType := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	
	transactions, total, err := h.platformAccountService.GetPlatformTransactions(transactionType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"total":        total,
		"page":         page,
		"page_size":    pageSize,
	})
}
