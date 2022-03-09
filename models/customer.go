package models

import (
	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp/v3"
	service "service_customer/service/rayan/proto"
)

type Customer struct {
	gorm.Model
	SejamReferenceCode string          `json:"sejam_reference_code" gorm:"size:100"`
	NormalNationalCode string          `json:"normal_national_code gorm:"size:10"`
	UserName           string          `json:"user_name" gorm:"size:100"`
	Password           string          `json:"user_name" gorm:"size:100"`
	IsActive           bool            `json:"is_active"`
	IsRayanService     bool            `json:"is_rayan_service"`
	PhonePerson        []PhonePerson   `json:"phone_persons"`
	PrivateInfo        CustomerPrivate `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	LegalInfo          CustomerLegal   `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	JobInfos           []JobInfo       `gorm:"foreignKey:CustomerId"`
}

func (c *Customer) TableName() string {
	return "customer_customer"
}

func (c Customer) SetBulkData(customers []*service.FundCustomerList) {
			for _, v := range customers {
				c.Set(*v)
			}
}

func (c Customer) Set(customer service.FundCustomerList) {
	pp.Print(customer)
}
