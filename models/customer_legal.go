package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerLegal struct {
	gorm.Model

	CompanyName            string `json:"company_name" gorm:"size:60"`
	RegisterNumber         string `json:"register_number" gorm:"size:60"`
	RegisterPlace          string `json:"register_place" gorm:"size:60"`
	RegisterDate           string `json:"register_date" gorm:"size:100"`
	EconomicCode           string `json:"economic_code" gorm:"size:100"`
	EvidenceReleaseCompany string `json:"evidence_release_company" gorm:"size:100"`
	EvidenceReleaseDate    string `json:"evidence_release_date" gorm:"size:100"`
	EvidenceExpirationDate string `json:"birth_date" gorm:"size:100"`
	CustomerId             uint   `json:"customer" `
}

func (cl *CustomerLegal) TableName() string {
	return "customer_legal_info"
}
