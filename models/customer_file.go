package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerFile struct {
	gorm.Model
	FileData   string `json:"file_data" gorm:"type:text"`
	CustomerId uint   `json:"customer" gorm:"Column:customer_id"`
	FileTypeId uint   `json:"file_type_id" gorm:"Column:file_type_id"`
}

func (cf *CustomerFile) TableName() string {
	return "customer_file"
}
