package models

import (
	"github.com/jinzhu/gorm"
)

type ComexVisitor struct {
	gorm.Model
	ComexIdRayan int32  `json:"comex_id_rayan" gorm:"Column:comex_id_rayan"`
	FullName     string `json:"full_name" gorm:"Column:full_name;size:60"`
	Rate         int32  `json:"rate" gorm:"Column:rate"`
	TypeMebbcoId uint   `gorm:"Column:type_mebbco_id"`
}

func (cv *ComexVisitor) TableName() string {
	return "customer_comex_visitor"
}
