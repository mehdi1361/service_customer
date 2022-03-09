package models

import (
	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	Name      string     `json:"name" gorm:"size:10;unique"`
	CountryId uint       `json="country_id" gorm:"Column:country_id"`
	Provinces []Province `json="provinces"`
}

func (p *Country) TableName() string {
	return "base_countries"
}
