package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Province struct {
	gorm.Model
	Name            string `json:"name" gorm:"size:10;unique"`
	ProvinceId      uint   `json="coun_id" gorm:"Column:province_id"`
	RayanProvinceId int    `json="province_id" gorm:"Column:rayan_province_id;unique"`
	CountryId       uint   `json="country_id" gorm:"Column:country_id"`
	Cities          []City `json="cities"`
}

func (p *Province) TableName() string {
	return "base_province"
}

func (p Province) GetOrCreate(data ProvinceParam) (*Province, error) {

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
	province := Province{}
	db.Find(&province, "rayan_province_id=?", data.ProvinceId)

	tx := db.Begin()

	if province.ID == 0 {
		newProvince := &Province{
			Name:            data.ProvinceName,
			RayanProvinceId: data.ProvinceId,
		}
		if err := tx.Create(newProvince).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		db.Find(&province, "rayan_province_id=?", data.ProvinceId)
	}

	err = tx.Commit().Error
	return &province, err
}
