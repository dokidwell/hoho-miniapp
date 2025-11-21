package handlers

import (
	"net/http"
	"strconv"

	"hoho-miniapp/backend/services"

	"github.com/gin-gonic/gin"
)

type AdminCreationHandler struct {
	creationService *services.CreationService
}

func NewAdminCreationHandler(creationService *services.CreationService) *AdminCreationHandler {
	return &AdminCreationHandler{
		creationService: creationService,
	}
}

// GetCreationList 获取创作列表（管理员）
func (h *AdminCreationHandler) GetCreationList(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	
	var creations interface{}
	var total int64
	var err error
	
	if status == "pending" {
		creations, total, err = h.creationService.GetPendingCreations(page, pageSize)
	} else {
		creations, total, err = h.creationService.GetPendingCreations(page, pageSize)
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

// GetCreationDetail 获取创作详情
func (h *AdminCreationHandler) GetCreationDetail(c *gin.Context) {
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	creation, err := h.creationService.GetCreationByID(uint(creationID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "创作不存在"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"creation": creation})
}

// ApproveCreation 批准创作
func (h *AdminCreationHandler) ApproveCreation(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := h.creationService.ApproveCreation(uint(creationID), adminID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "审核通过"})
}

// RejectCreation 拒绝创作
func (h *AdminCreationHandler) RejectCreation(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	creationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.creationService.RejectCreation(uint(creationID), adminID.(uint), req.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "已拒绝"})
}
