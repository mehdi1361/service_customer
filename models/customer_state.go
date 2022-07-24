package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerState struct {
	gorm.Model
	CustomerId uint `json:"customer_id" gorm:"Column:customer_id"`
	StateId    uint `json:"state_id" gorm:"Column:state_id"`
	Confirm    bool `json:"confirm" gorm:"Column:confirm"`
}

func (cs *CustomerState) TableName() string {
	return "customer_state"
}
