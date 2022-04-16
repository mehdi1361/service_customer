package models

import (
	"github.com/jinzhu/gorm"
	utils "service_customer/utils"
)

type VerificationCode struct {
	gorm.Model
	Code       string `json:"name" gorm:"size:4"`
	IsActive   bool   `json:"is_active" gorm:"Column:is_active;default:true"`
	CustomerId uint   `json:"customer_id"`
}

func (v *VerificationCode) TableName() string {
	return "customer_verification"
}

func (v VerificationCode) SendVerificationCode(customer Customer) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Model(&VerificationCode{}).Where("customer_id=?", customer.ID).Update("is_active", false)

	verificarionCode := VerificationCode{
		Code:       utils.RandomCodeGenerate(4),
		CustomerId: customer.ID,
		IsActive:   true,
	}
	db.Create(&verificarionCode)

	return nil
}
