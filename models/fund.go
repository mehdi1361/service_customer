package models

import (
	"github.com/jinzhu/gorm"
)

type Fund struct {
	gorm.Model
	isStockCreditPurchase bool   `json:"is_stock_credit_purchase" gorm:"Column:is_stock_credit_purchase"`
	CustomerStatusName    string `json:"customer_status_name" gorm:"Column:customer_status_name;size:60"`
	CustomerId            int64  `json:"customer_id" gorm:"Column:customer_id"`
	ReferredBy            string `json:"referred_by" gorm:"Column:referred_by;size:60"`
	BranchId              uint   `gorm:"Column:branch_id"`
	fundName              string `json:"fund_name" gorm:"Column:fund_name;size:60"`
}
func (f Fund) TableName() string {
	return "customer_fund"
}
