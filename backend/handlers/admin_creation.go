package handlers

import (
	"net/http"
	"strconv"

	"hoho-api/models"
	"hoho-api/services"

	"github.com/gin-gonic/gin"
)

// GetPendingCreations 获取待审核的创作列表
func GetPendingCreations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	creations, total, err := services.GetCreationsByStatus("pending", page, pageSize)
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

// ApproveCreation 批准创作
func ApproveCreation(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		TotalSupply int    `json:"total_supply"` // 管理员可修改发行数量
		Price       string `json:"price"`        // 管理员可修改价格
		Reason      string `json:"reason"`       // 审核意见
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.ApproveCreation(uint(creationID), adminID.(uint), req.TotalSupply, req.Price, req.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "审核通过"})
}

// RejectCreation 拒绝创作
func RejectCreation(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		Reason string `json:"reason" binding:"required"` // 拒绝原因
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := services.RejectCreation(uint(creationID), adminID.(uint), req.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "已拒绝"})
}

// PublishCreation 发布创作（上架销售）
func PublishCreation(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		PublishTime string `json:"publish_time"` // 发布时间（可选，为空则立即发布）
	}
	
	c.ShouldBindJSON(&req)
	
	if err := services.PublishCreation(uint(creationID), adminID.(uint), req.PublishTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "已发布"})
}

// GetAllCreations 获取所有创作列表（管理员）
func GetAllCreations(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	var creations []models.Creation
	var total int64
	var err error
	
	if status != "" {
		creations, total, err = services.GetCreationsByStatus(status, page, pageSize)
	} else {
		creations, total, err = services.GetAllCreations(page, pageSize)
	}
	
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
