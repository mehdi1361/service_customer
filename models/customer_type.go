package models

import (
	"github.com/jinzhu/gorm"
)

type FileType struct {
	gorm.Model
	Name          string         `json:"name" gorm:"Column:name;unique"`
	FaName        string         `json:"fa_name" gorm:"Column:fa_name"`
	IsActive      bool           `json:"is_active" gorm:"Column:is_active"`
	IsForce       bool           `json:"is_force" gorm:"Column:is_force"`
	CustomerFiles []CustomerFile `gorm:"foreignKey:FileTypeId"`
}

func (ft *FileType) TableName() string {
	return "base_file_type"
}
