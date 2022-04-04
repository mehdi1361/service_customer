package models

import (
	"github.com/jinzhu/gorm"
)

type Fund struct {
	gorm.Model
	isStockCreditPurchase bool   `json:"is_stock_credit_purchase" gorm:"Column:is_stock_credit_purchase"`
	CustomerId            int64  `json:"customer_id" gorm:"Column:customer_id"`
	ReferredBy            string `json:"referred_by" gorm:"Column:referred_by;size:60"`
	BranchId              uint   `gorm:"Column:branch_id"`
	FundName              string `json:"fund_name" gorm:"Column:fund_name;size:60;primaryKey"`
	Personality           string `json:"personality" gorm:"Column:personality;size:60"`
	FullName              string `json:"fullname" gorm:"Column:fullname;size:100"`
	IssunigCity           string `json:"issuing_city" gorm:"Column:issuing_city;size:100"`
	Nationality           string `json:"nationality" gorm:"Column:nationality;size:100"`
	NationalIdentifier    string `json:"national_identifier" gorm:"Column:national_identifier;size:100"`
	CreationDate          string `json:"creation_date" gorm:"Column:creation_date;size:100"`
	IsProfitIssue         string `json:"is_profit_issue" gorm:"Column:is_profit_issue;size:100"`
	Status                string `json:"status" gorm:"Column:status;size:100"`
	AccountNumber         string `json:"account_number" gorm:"Column:account_number;size:100"`
	CustomerServiceId     uint   `gorm:"Column:customer_service_id;primaryKey"`
}

func (f Fund) TableName() string {
	return "customer_fund"
}
