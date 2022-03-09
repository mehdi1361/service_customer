package models

import (
	"github.com/jinzhu/gorm"
)

type FinancialInfo struct {
	gorm.Model
	AssetValue             int32 `json:"asset_value" gorm:"Column:asset_value"`
	InCommingAverage       int32 `json:"incoming_average" gorm:"Column:asset_value"`
	SExchangeTransaction   int32 `json:"s_exchange_tranasction" gorm:"Column:s_exchange_tranasction"`
	CExchangeTransaction   int32 `json:"c_exchange_tranasction" gorm:"Column:c_exchange_tranasction"`
	OutExchangeTransaction int32 `json:"out_exchange_tranasction" gorm:"Column:out_exchange_tranasction"`
	TransactionLevel       int32 `json:"tranasction_level" gorm:"Column:tranasction_level"`
	TradingKnowledeLevel   int32 `json:"trading_knowledge_level" gorm:"Column:trading_knowledge_level"`
}

func (fi *FinancialInfo) TableName() string {
	return "customer_financial_info"
}
