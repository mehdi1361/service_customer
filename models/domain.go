package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Domain struct {
	gorm.Model
	Title      string `json:"title" gorm:"size:10;Null"`
	RayanId    int64  `json:"rayan_id" gorm: "Column:rayan_id"`
	CustomerId uint   `json:"customer" gorm:"Column:customer_id"`
}

func (d *Domain) TableName() string {
	return "base_mebbco_domain"
}


func (d Domain) GetOrCreate(data DomainParams) (*Domain, error) {

	db, err := Connect()
	if err != nil {
		return nil, err
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
	domain := Domain{}
	db.Find(&domain, "rayan_id=?", data.RayanId)

	tx := db.Begin()

	if domain.ID == 0 {

		if err := tx.Create(
			&Domain{
				RayanId:    data.RayanId,
				Title:      data.Title,
				CustomerId: data.CustomerId,
			},
		).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	return nil, err
}
