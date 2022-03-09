package models

import (
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	Title    string    `json:"title" gorm:"size:10;unique"`
	JobId    string    `json:"job_id" gorm:"size:10"`
	JobInfos []JobInfo `gorm:"foreignKey:JobId"`
}

func (mb *Job) TableName() string {
	return "base_job"
}
