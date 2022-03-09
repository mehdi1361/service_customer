package models

import (
	"github.com/jinzhu/gorm"
)

type BankBranch struct {
	gorm.Model
	Name         string        `json:"name" gorm:"size:10;unique"`
	BankId       uint          `gorm:"Column:bank_id"`
	BankBranchId uint          `gorm:"Column:bank_branch_id"`
	CityId       uint          `gorm:"Column:city_id"`
	BankAccount  []BankAccount `json:"bank_accounts"`
}

func (b *BankBranch) TableName() string {
	return "base_bank_branch"
}
