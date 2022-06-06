
package models

import (
	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:10;Null"`
	IdRayan int32  `json:"id_rayan"`
}

func (d *Domain) TableName() string {
	return "base_mebbco_domain"
}
