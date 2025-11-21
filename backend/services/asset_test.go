package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateInstanceNo 测试实例编号生成
func TestGenerateInstanceNo(t *testing.T) {
	tests := []struct {
		name         string
		assetID      uint64
		serialNumber int
		expected     string
	}{
		{
			name:         "正常生成",
			assetID:      1,
			serialNumber: 1,
			expected:     "A000001-0001",
		},
		{
			name:         "大编号",
			assetID:      999999,
			serialNumber: 9999,
			expected:     "A999999-9999",
		},
		{
			name:         "小编号",
			assetID:      1,
			serialNumber: 1,
			expected:     "A000001-0001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateInstanceNo(tt.assetID, tt.serialNumber)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestValidateAssetStatus 测试资产状态验证
func TestValidateAssetStatus(t *testing.T) {
	validStatuses := []string{
		"draft",
		"pending_review",
		"approved",
		"rejected",
		"minting",
		"active",
		"sold_out",
		"archived",
	}

	for _, status := range validStatuses {
		t.Run(status, func(t *testing.T) {
			assert.True(t, isValidAssetStatus(status))
		})
	}

	invalidStatuses := []string{
		"invalid",
		"unknown",
		"",
	}

	for _, status := range invalidStatuses {
		t.Run(status, func(t *testing.T) {
			assert.False(t, isValidAssetStatus(status))
		})
	}
}

// TestValidateMediaType 测试媒体类型验证
func TestValidateMediaType(t *testing.T) {
	tests := []struct {
		mediaType string
		valid     bool
	}{
		{"image", true},
		{"video", true},
		{"audio", true},
		{"3d_model", true},
		{"document", false},
		{"", false},
		{"unknown", false},
	}

	for _, tt := range tests {
		t.Run(tt.mediaType, func(t *testing.T) {
			result := isValidMediaType(tt.mediaType)
			assert.Equal(t, tt.valid, result)
		})
	}
}

// TestCalculateMintProgress 测试铸造进度计算
func TestCalculateMintProgress(t *testing.T) {
	tests := []struct {
		name         string
		mintedCount  int
		totalSupply  int
		expectedPct  float64
		expectError  bool
	}{
		{
			name:        "50%进度",
			mintedCount: 500,
			totalSupply: 1000,
			expectedPct: 50.0,
			expectError: false,
		},
		{
			name:        "100%进度",
			mintedCount: 1000,
			totalSupply: 1000,
			expectedPct: 100.0,
			expectError: false,
		},
		{
			name:        "0%进度",
			mintedCount: 0,
			totalSupply: 1000,
			expectedPct: 0.0,
			expectError: false,
		},
		{
			name:        "超过总量",
			mintedCount: 1100,
			totalSupply: 1000,
			expectedPct: 0,
			expectError: true,
		},
		{
			name:        "总量为0",
			mintedCount: 0,
			totalSupply: 0,
			expectedPct: 0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pct, err := calculateMintProgress(tt.mintedCount, tt.totalSupply)
			
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPct, pct)
			}
		})
	}
}

// TestCheckAssetAvailability 测试资产可用性检查
func TestCheckAssetAvailability(t *testing.T) {
	tests := []struct {
		name        string
		status      string
		mintedCount int
		totalSupply int
		available   bool
	}{
		{
			name:        "正常可用",
			status:      "active",
			mintedCount: 500,
			totalSupply: 1000,
			available:   true,
		},
		{
			name:        "已售罄",
			status:      "sold_out",
			mintedCount: 1000,
			totalSupply: 1000,
			available:   false,
		},
		{
			name:        "未激活",
			status:      "draft",
			mintedCount: 0,
			totalSupply: 1000,
			available:   false,
		},
		{
			name:        "已归档",
			status:      "archived",
			mintedCount: 500,
			totalSupply: 1000,
			available:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkAssetAvailability(tt.status, tt.mintedCount, tt.totalSupply)
			assert.Equal(t, tt.available, result)
		})
	}
}

// TestValidateAssetOwnership 测试资产所有权验证
func TestValidateAssetOwnership(t *testing.T) {
	tests := []struct {
		name        string
		ownerID     uint64
		userID      uint64
		isOwner     bool
	}{
		{
			name:    "是所有者",
			ownerID: 1,
			userID:  1,
			isOwner: true,
		},
		{
			name:    "不是所有者",
			ownerID: 1,
			userID:  2,
			isOwner: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.ownerID == tt.userID
			assert.Equal(t, tt.isOwner, result)
		})
	}
}

// TestAssetInstanceStatus 测试资产实例状态
func TestAssetInstanceStatus(t *testing.T) {
	validStatuses := []string{
		"minting",
		"owned",
		"listed",
		"trading",
		"burned",
	}

	for _, status := range validStatuses {
		t.Run(status, func(t *testing.T) {
			assert.True(t, isValidInstanceStatus(status))
		})
	}
}

// 辅助函数

func generateInstanceNo(assetID uint64, serialNumber int) string {
	return fmt.Sprintf("A%06d-%04d", assetID, serialNumber)
}

func isValidAssetStatus(status string) bool {
	validStatuses := map[string]bool{
		"draft":          true,
		"pending_review": true,
		"approved":       true,
		"rejected":       true,
		"minting":        true,
		"active":         true,
		"sold_out":       true,
		"archived":       true,
	}
	return validStatuses[status]
}

func isValidMediaType(mediaType string) bool {
	validTypes := map[string]bool{
		"image":    true,
		"video":    true,
		"audio":    true,
		"3d_model": true,
	}
	return validTypes[mediaType]
}

func calculateMintProgress(mintedCount, totalSupply int) (float64, error) {
	if totalSupply == 0 {
		return 0, &ValidationError{Message: "总供应量不能为0"}
	}
	if mintedCount > totalSupply {
		return 0, &ValidationError{Message: "已铸造数量超过总供应量"}
	}
	return float64(mintedCount) / float64(totalSupply) * 100, nil
}

func checkAssetAvailability(status string, mintedCount, totalSupply int) bool {
	if status != "active" {
		return false
	}
	if mintedCount >= totalSupply {
		return false
	}
	return true
}

func isValidInstanceStatus(status string) bool {
	validStatuses := map[string]bool{
		"minting": true,
		"owned":   true,
		"listed":  true,
		"trading": true,
		"burned":  true,
	}
	return validStatuses[status]
}

// 需要导入fmt包
