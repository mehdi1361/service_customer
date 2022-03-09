package models

import (
	"github.com/jinzhu/gorm"
)

type PhonePerson struct {
	gorm.Model
	AccountNumber string `json:"account_number gorm:"size:60"`
	Sheba         string `json:"sheba" gorm:"size:60"`
	IsActive      bool   `json:"is_active" gorm:"Column:is_active"`
	CustomerId    uint   `json:"customer_id" gorm:"Column:customer_id"`
}

func (pp *PhonePerson) TableName() string {
	return "customer_phone_person"
}
