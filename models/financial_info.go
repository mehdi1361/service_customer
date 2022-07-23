package models

import (
	"github.com/jinzhu/gorm"
)

type FinancialInfo struct {
	gorm.Model
	AssetsValue            int64  `json:"asset_value" gorm:"Column:asset_value"`
	InComingAverage        int64  `json:"incoming_average" gorm:"Column:incoming_average"`
	SExchangeTransaction   int64  `json:"s_exchange_tranasction" gorm:"Column:s_exchange_tranasction"`
	CExchangeTransaction   int64  `json:"c_exchange_tranasction" gorm:"Column:c_exchange_tranasction"`
	OutExchangeTransaction int64  `json:"out_exchange_tranasction" gorm:"Column:out_exchange_tranasction"`
	TransactionLevel       string `json:"tranasction_level" gorm:"Column:tranasction_level"`
	TradingKnowledgeLevel  string `json:"trading_knowledge_level" gorm:"Column:trading_knowledge_level"`
	CompanyPurpose         string `json:"company_purpose" gorm:"Column:company_purpose"`
	ReferenceRateCompany   string `json:"reference_rate_company" gorm:"Column:reference_rate_company"`
	RateDate               string `json:"rate_date" gorm:"Column:rate_date"`
	Rate                   string `json:"rate" gorm:"Column:rate"`
	CustomerId             uint   `json:"customer" gorm:"unique"`
}

func (fi *FinancialInfo) TableName() string {
	return "customer_financial_info"
}

func (j FinancialInfo) GetOrCreate(d FinancialInfo) (*FinancialInfo, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	finInfo := FinancialInfo{}
	db.Find(&finInfo, "customer_id=?", d.CustomerId)

	if finInfo.ID == 0 {
		finInfo = FinancialInfo{
			AssetsValue:            d.AssetsValue,
			InComingAverage:        d.InComingAverage,
			SExchangeTransaction:   d.SExchangeTransaction,
			CExchangeTransaction:   d.CExchangeTransaction,
			OutExchangeTransaction: d.OutExchangeTransaction,
			TransactionLevel:       d.TransactionLevel,
			TradingKnowledgeLevel:  d.TradingKnowledgeLevel,
			CompanyPurpose:         d.CompanyPurpose,
			ReferenceRateCompany:   d.ReferenceRateCompany,
			RateDate:               d.RateDate,
			Rate:                   d.Rate,
		}

		db.Create(&finInfo)
	}
	return &finInfo, nil

}
