package models

import (
	"github.com/jinzhu/gorm"
)

type MebbcoBranch struct {
	gorm.Model
	Title         string         `json:"title" gorm:"size:10;unique"`
	TypeMebbco    string         `json:"type_mebbco" gorm:"size:50;column:type_mebbco"`
	IdRayan       int32          `json:"id_rayan"`
	Portfos       []Portfo       `gorm:"foreignKey:TypeMebbcoId"`
	ComexVisitors []ComexVisitor `gorm:"foreignKey:TypeMebbcoId"`
	TradingCodes  []TradingCode  `gorm:"foreignKey:TypeMebbcoId"`
}

func (mb *MebbcoBranch) TableName() string {
	return "base_mebbco_branch"
}
