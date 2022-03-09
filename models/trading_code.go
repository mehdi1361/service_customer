package models

import (
	"github.com/jinzhu/gorm"
)

type TradingCode struct {
	gorm.Model
	Type       string `json:"Type" gorm:"size:50;Column:type"`
	FirstPart  string `json:"first_part" gorm:"size:70;Column:first_part"`
	SecondPart string `json:"second_part" gorm:"size:70;Column:second_part"`
	Code       string `json:"code" gorm:"size:70;Column:code"`
	IsDefault  bool   `json:"is_default" gorm:"Column:is_default"`
}

func (tc *TradingCode) TableName() string {
	return "base_job"
}
