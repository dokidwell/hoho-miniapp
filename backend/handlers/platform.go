package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetPlatformAccount 获取平台账户信息（阳光账户）
func GetPlatformAccount(c *gin.Context) {
	account, err := services.GetPlatformAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"account": account})
}

// GetPlatformTransactions 获取平台交易记录
func GetPlatformTransactions(c *gin.Context) {
	transactionType := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	
	transactions, total, err := services.GetPlatformTransactions(transactionType, page, pageSize)
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
