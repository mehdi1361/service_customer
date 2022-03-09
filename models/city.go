package models

import (
	"github.com/jinzhu/gorm"
)

type City struct {
	gorm.Model
	Name       string       `json:"name" gorm:"size:10;unique"`
	ProvinceId uint         `json="province_id" gorm:"Column:province_id"`
	CityId     uint         `json="city_id" gorm:"Column:city_id"`
	BankBranch []BankBranch `json:"bank_branch"`
}

func (c *City) TableName() string {
	return "base_cities"
}
