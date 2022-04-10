package models

import (
	"github.com/jinzhu/gorm"
)

type Bank struct {
	gorm.Model
	Title      string       `json:"title" gorm:"size:10"`
	BankId     uint         `json="bank_id" gorm:"Column:bank_id;unique"`
	BankBranch []BankBranch `json:"bank_branch"`
}

func (b *Bank) TableName() string {
	return "base_bank"
}

