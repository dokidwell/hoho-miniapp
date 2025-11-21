package services

import (
	"hoho-api/database"
	"hoho-api/models"
	
	"github.com/shopspring/decimal"
)

// GetPlatformAccount 获取平台账户信息
func GetPlatformAccount() (*models.PlatformAccount, error) {
	var account models.PlatformAccount
	if err := database.DB.First(&account, 1).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// RecordPlatformIncome 记录平台收入
func RecordPlatformIncome(incomeType, amount, description string, relatedID *uint) error {
	// 获取平台账户
	var account models.PlatformAccount
	if err := database.DB.First(&account, 1).Error; err != nil {
		return err
	}
	
	// 计算新余额
	currentBalance, _ := decimal.NewFromString(account.TotalBalance)
	amountDecimal, _ := decimal.NewFromString(amount)
	newBalance := currentBalance.Add(amountDecimal)
	
	// 更新账户余额
	updates := map[string]interface{}{
		"total_balance": newBalance.String(),
	}
	
	if incomeType == "commission" {
		commissionIncome, _ := decimal.NewFromString(account.CommissionIncome)
		updates["commission_income"] = commissionIncome.Add(amountDecimal).String()
	} else if incomeType == "fee" {
		feeIncome, _ := decimal.NewFromString(account.FeeIncome)
		updates["fee_income"] = feeIncome.Add(amountDecimal).String()
	}
	
	if err := database.DB.Model(&account).Updates(updates).Error; err != nil {
		return err
	}
	
	// 记录交易
	transaction := &models.PlatformTransaction{
		Type:         incomeType,
		Amount:       amount,
		BalanceAfter: newBalance.String(),
		Description:  description,
		RelatedID:    relatedID,
	}
	
	return database.DB.Create(transaction).Error
}

// RecordPlatformExpense 记录平台支出
func RecordPlatformExpense(amount, description string, relatedID *uint) error {
	// 获取平台账户
	var account models.PlatformAccount
	if err := database.DB.First(&account, 1).Error; err != nil {
		return err
	}
	
	// 计算新余额
	currentBalance, _ := decimal.NewFromString(account.TotalBalance)
	amountDecimal, _ := decimal.NewFromString(amount)
	newBalance := currentBalance.Sub(amountDecimal)
	
	// 更新账户余额
	totalExpense, _ := decimal.NewFromString(account.TotalExpense)
	updates := map[string]interface{}{
		"total_balance":  newBalance.String(),
		"total_expense":  totalExpense.Add(amountDecimal).String(),
	}
	
	if err := database.DB.Model(&account).Updates(updates).Error; err != nil {
		return err
	}
	
	// 记录交易
	transaction := &models.PlatformTransaction{
		Type:         "expense",
		Amount:       amount,
		BalanceAfter: newBalance.String(),
		Description:  description,
		RelatedID:    relatedID,
	}
	
	return database.DB.Create(transaction).Error
}

// GetPlatformTransactions 获取平台交易记录
func GetPlatformTransactions(transactionType string, page, pageSize int) ([]models.PlatformTransaction, int64, error) {
	var transactions []models.PlatformTransaction
	var total int64
	
	query := database.DB.Model(&models.PlatformTransaction{})
	
	if transactionType != "" {
		query = query.Where("type = ?", transactionType)
	}
	
	query.Count(&total)
	
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&transactions).Error; err != nil {
		return nil, 0, err
	}
	
	return transactions, total, nil
}
