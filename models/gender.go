package models

import (
	"github.com/jinzhu/gorm"
)

type Gender struct {
	gorm.Model
	Name            string `json:"name" gorm:"size:10;unique"`
	CustomerPrivate []CustomerPrivate
}

func (g *Gender) TableName() string {
	return "base_gender"
}
