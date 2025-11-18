package handlers

import (
	"hoho-miniapp/backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AssetHandler 定义藏品相关的HTTP处理函数
type AssetHandler struct {
	AssetService *services.AssetService
}

// NewAssetHandler 创建一个新的AssetHandler实例
func NewAssetHandler(assetService *services.AssetService) *AssetHandler {
	return &AssetHandler{AssetService: assetService}
}

// SubmitMintRequest 提交铸造请求
func (h *AssetHandler) SubmitMintRequest(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		CollectionID uint64 `json:"collection_id" binding:"required"`
		Name         string `json:"name" binding:"required"`
		Description  string `json:"description"`
		MediaURL     string `json:"media_url" binding:"required"`
		MediaType    string `json:"media_type" binding:"required,oneof=image video audio"`
		TotalSupply  int    `json:"total_supply" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "details": err.Error()})
		return
	}

	asset, err := h.AssetService.SubmitMintRequest(userID.(uint64), req.CollectionID, req.Name, req.Description, req.MediaURL, req.MediaType, req.TotalSupply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "提交铸造请求失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "铸造请求已提交，等待审核",
		"data": asset,
	})
}

// ListAssets 处理获取藏品列表请求
func (h *AssetHandler) ListAssets(c *gin.Context) {
	// TODO: 实现分页、筛选、排序逻辑
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "藏品列表 - TODO"})
}

// GetAssetDetail 处理获取藏品详情请求
func (h *AssetHandler) GetAssetDetail(c *gin.Context) {
	assetIDStr := c.Param("id")
	assetID, err := strconv.ParseUint(assetIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "藏品ID格式错误"})
		return
	}

	asset, err := h.AssetService.GetAssetByID(assetID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "藏品不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": asset,
	})
}
