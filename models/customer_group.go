package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerGroup struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:10;unique"`
	IdRayan int32  `json:"id_rayan"`
}

func (cg *CustomerGroup) TableName() string {
	return "base_mebbco_customer_group"
}
