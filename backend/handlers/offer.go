package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// CreateOffer 创建出价
func CreateOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var req struct {
		ArtworkInstanceID uint   `json:"artwork_instance_id" binding:"required"`
		Price             string `json:"price" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	offer, err := services.CreateOffer(userID.(uint), req.ArtworkInstanceID, req.Price)
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
func GetMyOffers(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	offers, total, err := services.GetUserOffers(userID.(uint), status, page, pageSize)
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

// GetReceivedOffers 获取收到的出价列表
func GetReceivedOffers(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	offers, total, err := services.GetReceivedOffers(userID.(uint), status, page, pageSize)
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

// AcceptOffer 接受出价
func AcceptOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	offerID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.AcceptOffer(uint(offerID), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "出价已接受"})
}

// RejectOffer 拒绝出价
func RejectOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	offerID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.RejectOffer(uint(offerID), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "出价已拒绝"})
}

// CancelOffer 取消出价
func CancelOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	offerID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := services.CancelOffer(uint(offerID), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "出价已取消"})
}
