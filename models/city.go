package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type City struct {
	gorm.Model
	Name        string       `json:"name" gorm:"size:10;unique"`
	ProvinceId  uint         `json="province_id" gorm:"Column:province_id"`
	CityId      uint         `json="city_id" gorm:"Column:city_id"`
	RayanCityId int          `json="rayan_city_id" gorm:"Column:rayan_city_id"`
	BankBranch  []BankBranch `json:"bank_branch"`
}

func (c *City) TableName() string {
	return "base_cities"
}

func (c City) GetOrCreate(data SetCityParams) (*City, error) {
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
	city := City{}
	db.Find(&city, "rayan_city_id=?", data.CityId)

	tx := db.Begin()

	if city.ID == 0 {

		province, err := Province{}.GetOrCreate(
			ProvinceParam{
				ProvinceId:   data.ProvinceId,
				ProvinceName: data.ProvinceName,
			},
		)

		if err != nil {
			return nil, err
		}

		if err := tx.Create(
			&City{
				Name:        data.CityName,
				RayanCityId: data.CityId,
				ProvinceId:  province.ID,
			},
		).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	return nil, err
}
