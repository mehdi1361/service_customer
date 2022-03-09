package models

import (
	"github.com/jinzhu/gorm"
)

type BankAccount struct {
	gorm.Model
	AccountNumber string `json:"account_number gorm:"size:60"`
	Sheba         string `json:"sheba" gorm:"size:60"`
	IsDefault     bool   `json:"is_default" gorm:"Column:is_default"`
	BranchId      uint   `json:"branch_id" gorm:"Column:branch_id"`
}

func (ba *BankAccount) TableName() string {
	return "customer_bank_account"
}
