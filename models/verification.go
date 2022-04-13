package models

import (
	"github.com/jinzhu/gorm"
)

type VerificationCode struct {
	gorm.Model
	Code     string `json:"name" gorm:"size:4"`
	IsActive bool   `json:"is_active" gorm:"Column:is_active;default:true"`

	Customer_id uint `json:"customer_id"`
}

func (v *VerificationCode) TableName() string {
	return "customer_verification"
}
