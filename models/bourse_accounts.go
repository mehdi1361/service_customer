package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BourseAccounts struct {
	gorm.Model
	RayanId    int64    `json:"rayan_id" gorm:"Column:rayan_id"`
	Name       string `json:"name" gorm:"size:100"`
	Number     string `json:"number" gorm:"size:100;Column:number"`
	IsDefault  int64    `json:"is_default" gorm:"Column:is_default"`
	CustomerId uint   `json:"customer_id" gorm:"Column:customer_id"`
}

func (c *BourseAccounts) TableName() string {
	return "customer_broker_bourse_accounts"
}

func (ba BourseAccounts) GetOrCreate(data BourseAccountParam) (*BourseAccounts, error) {
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
	bourseAccounts := BourseAccounts{}
	db.Find(&bourseAccounts, "rayan_id=?", data.Id)

	tx := db.Begin()

	if bourseAccounts.ID == 0 {

		if err := tx.Create(
			&BourseAccounts{
				RayanId:    data.Id,
				Name:       data.Name,
				Number:     data.Number,
				IsDefault:  data.IsDefault,
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
