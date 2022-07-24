package models

import (
	"github.com/jinzhu/gorm"
)

type State struct {
	gorm.Model
	StateName string `json:"state_name" gorm:"size:100;Column:state_name"`
	Title     string `json:"title" gorm:"size:100"`
	IconClass string `json:"icon_class" gorm:"size:100;Column:icon_class"`
	StateId   int64  `json:"state_id" gorm:Column:state_id"`
}

func (s *State) TableName() string {
	return "base_state"
}
