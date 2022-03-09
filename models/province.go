package models

import (
	"github.com/jinzhu/gorm"
)

type Province struct {
	gorm.Model
	Name       string `json:"name" gorm:"size:10;unique"`
	ProvinceId uint   `json="coun_id" gorm:"Column:country_id"`
	CountryId  uint   `json="country_id" gorm:"Column:country_id"`
	Cities     []City `json="cities"`
}

func (p *Province) TableName() string {
	return "base_province"
}
