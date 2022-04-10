package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PhonePerson struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" gorm:"size:60;Column:phone_number" sql:"unique_index:idx_phone_customer`
	IsActive    bool   `json:"is_active" gorm:"Column:is_active"`
	IsMobile    bool   `json:"is_mobile" gorm:"Column:is_mobile"`
	MebbcoType  string `json:"mebbco_type gorm:"size:60;Column:mebbco_type"`
	CustomerId  uint   `json:"customer_id" gorm:"Column:customer_id" sql:"unique_index:idx_phone_customer`
}

func (pp *PhonePerson) TableName() string {
	return "customer_phone_person"
}

func (pp PhonePerson) SetPhone(phonePerson *PhonePerson) error {

	db, err := Connect()
	if err != nil {
		return err
	}

	sqlDB := db.DB()
	for {
		if e := sqlDB.Ping(); e == nil {
			break
		}
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(300)

	sqlDB.SetConnMaxLifetime(time.Hour)

	defer db.Close()
	phone := PhonePerson{}
	db.Find(&phone, "phone_number=? and customer_id=?", phonePerson.PhoneNumber, phonePerson.CustomerId)

	tx := db.Begin()

	if phone.ID == 0 {
		if err := tx.Create(&phonePerson).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}
