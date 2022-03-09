package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerPrivate struct {
	gorm.Model

	FirstName     string `json:"first_name" gorm:"size:60"`
	LastName      string `json:"last_name" gorm:"size:60"`
	FatherName    string `json:"father_name" gorm:"size:60"`
	SeriShChar    string `json:"seri_sh_char" gorm:"size:100"`
	SeriSh        string `json:"seri_sh" gorm:"size:100"`
	Serial        string `json:"serial" gorm:"size:100"`
	ShNumber      string `json:"sh_number" gorm:"size:100"`
	BirthDate     string `json:"birth_date" gorm:"size:100"`
	PlaceOfIssue  string `json:"place_of_issue" gorm:"size:100"`
	SignatureFile string `json:"signature_file" gorm:"type:text"`
	CustomerId    uint   `json:"customer" `
}

func (cp *CustomerPrivate) TableName() string {
	return "customer_private_info"
}
