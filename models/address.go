package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CustomerAddress struct {
	gorm.Model
	Address    string `json:"address" gorm:"Column:address;size:250"`
	PostalCode string `json:"postal_code" gorm:"Column:postal_code;size:20;Null"`
	CustomerId uint   `json="customer_id" gorm:"Column:customer_id"`
	MebbcoType string `json:"mebbco_type gorm:"size:60;Column:mebbco_type"`
}

func (ca *CustomerAddress) TableName() string {
	return "customer_address"
}

func (ca CustomerAddress) Set(data AddressParam) error {
	db, err := Connect()
	if err != nil {
		return  err
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
	customerAddress := CustomerAddress{}
	db.Find(&customerAddress, "address=? or postal_code=?", data.Address, data.PostalCode)

	tx := db.Begin()

	if customerAddress.ID == 0 {


		if err := tx.Create(
			&CustomerAddress{
				Address:        data.Address,
				PostalCode: data.PostalCode,
				CustomerId:  data.CustomerId,
				MebbcoType:  data.MebbcoType,
			},
		).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	return  err
}
