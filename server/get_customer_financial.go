package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) GetPersonFinancialInfo(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.PersonFinancialResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customer := models.Customer{}
	financialInfo := models.FinancialInfo{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&financialInfo, "customer_id=?", customer.ID)

	return &service.PersonFinancialResponse{
		AssetsValue:            financialInfo.AssetsValue,
		InComingAverage:        financialInfo.InComingAverage,
		SExchangeTransaction:   financialInfo.SExchangeTransaction,
		CExchangeTransaction:   financialInfo.CExchangeTransaction,
		OutExchangeTransaction: financialInfo.OutExchangeTransaction,
		TransactionLevel:       financialInfo.TransactionLevel,
		TradingKnowledgeLevel:  financialInfo.TradingKnowledgeLevel,
		CompanyPurpose:         financialInfo.CompanyPurpose,
		ReferenceRateCompany:   financialInfo.ReferenceRateCompany,
		RateDate:               financialInfo.RateDate,
		Rate:                   financialInfo.Rate,
	}, nil
}
