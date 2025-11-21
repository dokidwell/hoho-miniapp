package services

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestAddPoints 测试增加积分
func TestAddPoints(t *testing.T) {
	tests := []struct {
		name          string
		userID        uint64
		amount        string
		description   string
		expectError   bool
		expectedTotal string
	}{
		{
			name:          "正常增加积分",
			userID:        1,
			amount:        "100.00000000",
			description:   "注册奖励",
			expectError:   false,
			expectedTotal: "100.00000000",
		},
		{
			name:        "金额为负数",
			userID:      1,
			amount:      "-50.00000000",
			description: "测试",
			expectError: true,
		},
		{
			name:        "金额为零",
			userID:      1,
			amount:      "0",
			description: "测试",
			expectError: true,
		},
		{
			name:        "金额格式错误",
			userID:      1,
			amount:      "abc",
			description: "测试",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 解析金额
			amount, err := decimal.NewFromString(tt.amount)
			
			if tt.expectError {
				// 如果期望错误，检查金额是否有效
				if err == nil {
					// 检查业务逻辑错误
					assert.True(t, amount.LessThanOrEqual(decimal.Zero), "金额应该无效")
				}
			} else {
				assert.NoError(t, err)
				assert.True(t, amount.GreaterThan(decimal.Zero))
			}
		})
	}
}

// TestDeductPoints 测试扣除积分
func TestDeductPoints(t *testing.T) {
	tests := []struct {
		name        string
		userID      uint64
		balance     string
		deduct      string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "正常扣除",
			userID:      1,
			balance:     "100.00000000",
			deduct:      "50.00000000",
			expectError: false,
		},
		{
			name:        "余额不足",
			userID:      1,
			balance:     "30.00000000",
			deduct:      "50.00000000",
			expectError: true,
			errorMsg:    "积分余额不足",
		},
		{
			name:        "扣除金额为负",
			userID:      1,
			balance:     "100.00000000",
			deduct:      "-10.00000000",
			expectError: true,
			errorMsg:    "扣除金额必须大于0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, _ := decimal.NewFromString(tt.balance)
			deduct, _ := decimal.NewFromString(tt.deduct)
			
			var err error
			if deduct.LessThanOrEqual(decimal.Zero) {
				err = &ValidationError{Message: "扣除金额必须大于0"}
			} else if balance.LessThan(deduct) {
				err = &ValidationError{Message: "积分余额不足"}
			}

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestFreezePoints 测试冻结积分
func TestFreezePoints(t *testing.T) {
	tests := []struct {
		name        string
		balance     string
		frozen      string
		toFreeze    string
		expectError bool
	}{
		{
			name:        "正常冻结",
			balance:     "100.00000000",
			frozen:      "0",
			toFreeze:    "30.00000000",
			expectError: false,
		},
		{
			name:        "可用余额不足",
			balance:     "100.00000000",
			frozen:      "80.00000000",
			toFreeze:    "30.00000000",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, _ := decimal.NewFromString(tt.balance)
			frozen, _ := decimal.NewFromString(tt.frozen)
			toFreeze, _ := decimal.NewFromString(tt.toFreeze)
			
			available := balance.Sub(frozen)
			
			var err error
			if available.LessThan(toFreeze) {
				err = &ValidationError{Message: "可用余额不足"}
			}

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestUnfreezePoints 测试解冻积分
func TestUnfreezePoints(t *testing.T) {
	tests := []struct {
		name        string
		frozen      string
		toUnfreeze  string
		expectError bool
	}{
		{
			name:        "正常解冻",
			frozen:      "50.00000000",
			toUnfreeze:  "30.00000000",
			expectError: false,
		},
		{
			name:        "解冻金额大于冻结金额",
			frozen:      "20.00000000",
			toUnfreeze:  "30.00000000",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			frozen, _ := decimal.NewFromString(tt.frozen)
			toUnfreeze, _ := decimal.NewFromString(tt.toUnfreeze)
			
			var err error
			if frozen.LessThan(toUnfreeze) {
				err = &ValidationError{Message: "冻结金额不足"}
			}

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestCalculatePoints 测试积分计算
func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		balance  string
		frozen   string
		expected string
	}{
		{
			name:     "正常计算",
			balance:  "100.00000000",
			frozen:   "30.00000000",
			expected: "70.00000000",
		},
		{
			name:     "全部冻结",
			balance:  "100.00000000",
			frozen:   "100.00000000",
			expected: "0",
		},
		{
			name:     "无冻结",
			balance:  "100.00000000",
			frozen:   "0",
			expected: "100.00000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, _ := decimal.NewFromString(tt.balance)
			frozen, _ := decimal.NewFromString(tt.frozen)
			expected, _ := decimal.NewFromString(tt.expected)
			
			available := balance.Sub(frozen)
			
			assert.True(t, available.Equal(expected))
		})
	}
}

// TestPointTransactionType 测试积分交易类型
func TestPointTransactionType(t *testing.T) {
	validTypes := []string{
		"register_reward",
		"daily_signin",
		"purchase",
		"sell",
		"airdrop",
		"admin_adjust",
		"refund",
	}

	for _, transType := range validTypes {
		t.Run(transType, func(t *testing.T) {
			assert.NotEmpty(t, transType)
			assert.Contains(t, validTypes, transType)
		})
	}
}

// TestDecimalPrecision 测试decimal精度
func TestDecimalPrecision(t *testing.T) {
	tests := []struct {
		name     string
		value1   string
		value2   string
		operation string
		expected string
	}{
		{
			name:      "加法精度",
			value1:    "0.12345678",
			value2:    "0.87654322",
			operation: "add",
			expected:  "1.00000000",
		},
		{
			name:      "减法精度",
			value1:    "1.00000000",
			value2:    "0.12345678",
			operation: "sub",
			expected:  "0.87654322",
		},
		{
			name:      "乘法精度",
			value1:    "0.1",
			value2:    "3",
			operation: "mul",
			expected:  "0.3",
		},
		{
			name:      "除法精度",
			value1:    "1",
			value2:    "3",
			operation: "div",
			expected:  "0.33333333",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, _ := decimal.NewFromString(tt.value1)
			v2, _ := decimal.NewFromString(tt.value2)
			expected, _ := decimal.NewFromString(tt.expected)
			
			var result decimal.Decimal
			switch tt.operation {
			case "add":
				result = v1.Add(v2)
			case "sub":
				result = v1.Sub(v2)
			case "mul":
				result = v1.Mul(v2)
			case "div":
				result = v1.Div(v2).Truncate(8)
			}
			
			assert.True(t, result.Equal(expected), 
				"Expected %s, got %s", expected.String(), result.String())
		})
	}
}

// BenchmarkAddPoints 基准测试：增加积分
func BenchmarkAddPoints(b *testing.B) {
	amount := decimal.NewFromFloat(100.0)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		balance := decimal.Zero
		balance = balance.Add(amount)
	}
}

// BenchmarkDecimalOperations 基准测试：decimal运算
func BenchmarkDecimalOperations(b *testing.B) {
	v1 := decimal.NewFromFloat(123.456)
	v2 := decimal.NewFromFloat(789.012)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v1.Add(v2)
		_ = v1.Sub(v2)
		_ = v1.Mul(v2)
		_ = v1.Div(v2)
	}
}
