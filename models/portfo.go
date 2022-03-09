package models

import (
	"github.com/jinzhu/gorm"
)

type Portfo struct {
	gorm.Model
	CustomerId           int32  `json:"customer_id" gorm:"Column:customer_id"`
	DomainIdsId          int32  `json:"domain_id" gorm:"Column:domain_id"`
	IntroDate            string `json:"intro_date" gorm:"Column:intro_date;size:60"`
	GroupIdsId           int32  `json:"group_id" gorm:"Column:group_id"`
	Introducer           string `json:"out_exchange_tranasction" gorm:"Column:out_exchange_tranasction;size:60"`
	BourseAccountNumber  string `json:"bourse_account_number" gorm:"Column:bourse_account_number;size:60"`
	Comments             string `json:"comment" gorm:"Column:comment;size:250"`
	CalcMinimumCommision int32  `json:"c_exchange_tranasction" gorm:"Column:c_exchange_tranasction"`
	StockCreditPurchase  int32  `json:"stock_credit_purchase" gorm:"Column:stock_credit_purchase"`
	ParentNationalCode   string `json:"parent_national_code" gorm:"Column:parent_national_code;size:250"`
	ParentIsPortfo       bool   `json:"parent_is_portfo" gorm:"Column:parent_is_portfo;size:250"`
	ComexVisitorId       int32  `json:"comex_visitort_id" gorm:"Column:comex_visitor_id;size:250"`
	IsOnline             bool   `json:"is_online" gorm:"Column:is_online"`
	CustomerStatusId     int32  `json:"customer_status_id" gorm:"Column:customer_status_id"`
	EorderStatusId       int32  `json:"eorder_status_id" gorm:"Column:eorder_status_id"`
	SmsTransacrion       string `json:"sms_transaction" gorm:"Column:sms_transaction;size:50"`
	IsEPaymentCustomer   bool   `json:"is_e_payment_customer" gorm:"Column:is_e_payment_customer"`
	DlNumber             int32  `json:"dl_number" gorm:"Column:dl_number"`
	TypeMebbcoId         uint   `gorm:"Column:type_mebbco_id"`
}

func (po Portfo) TableName() string {
	return "customer_portfo"
}
