package server

import (
	"fmt"
	//	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) SejamRegisterPrivatePerson(ctx context.Context, e *service.SejamRegisterPrivatePersonRequest) (*service.SejamRegisterPrivatePersonResponse, error) {

	result := service.SejamRegisterPrivatePersonResponse{
		Id:      1,
		Message: "hello world",
	}

	isSejami := false

	if e.Profile.Status == "Sejami" {
		isSejami = true
	}

	customer, err := models.Customer{}.GetOrCreate(models.Customer{
		NormalNationalCode: e.Profile.NationalCode,
		IsSejami:           isSejami,
		SejamType:          e.Profile.CustomerSejamType,
	})

	if err != nil {
		return &service.SejamRegisterPrivatePersonResponse{
			Id:      400,
			Message: fmt.Sprintf("error: %s", err),
		}, nil
	}

	_, err = models.JobInfo{}.GetOrCreate(models.JobInfo{
		EmploymentDate:    e.JobInfo.EmploymentDate,
		CompanyName:       e.JobInfo.CompanyName,
		CompanyAddress:    e.JobInfo.CompanyAddress,
		CompanyPostalCode: e.JobInfo.CompanyPostalCode,
		CompanyEmail:      e.JobInfo.CompanyEmail,
		CompanyWebSite:    e.JobInfo.CompanyWebSite,
		CompanyCityPrefix: e.JobInfo.CompanyCityPrefix,
		CompanyPhone:      e.JobInfo.CompanyPhone,
		Position:          e.JobInfo.Position,
		CompanyFaxPrefix:  e.JobInfo.CompanyFaxPrefix,
		CompanyFax:        e.JobInfo.CompanyFax,
		JobId:             uint(e.JobInfo.JobId),
		JobTitle:          e.JobInfo.JobTitle,
		JobDescription:    e.JobInfo.JobDescription,
		CustomerId:        customer.ID,
	})

	_, err = models.CustomerPrivate{}.GetOrCreate(models.CustomerPrivate{
		FirstName:     e.PrivatePerson.FirstName,
		LastName:      e.PrivatePerson.LastName,
		Gender:        e.PrivatePerson.Gender,
		SeriSh:        e.PrivatePerson.SeriSh,
		Serial:        e.PrivatePerson.Serial,
		ShNumber:      e.PrivatePerson.ShNumber,
		BirthDate:     e.PrivatePerson.BirthDate,
		PlaceOfIssue:  e.PrivatePerson.PlaceOfIssue,
		PlaceOfBirth:  e.PrivatePerson.PlaceOfBirth,
		SignatureFile: e.PrivatePerson.SignatureFile,
		CustomerId:    customer.ID,
	})

	_, err = models.FinancialInfo{}.GetOrCreate(models.FinancialInfo{
		AssetsValue:            e.FinancialInfo.AssetsValue,
		InComingAverage:        e.FinancialInfo.InComingAverage,
		SExchangeTransaction:   e.FinancialInfo.SExchangeTransaction,
		CExchangeTransaction:   e.FinancialInfo.CExchangeTransaction,
		OutExchangeTransaction: e.FinancialInfo.OutExchangeTransaction,
		TransactionLevel:       e.FinancialInfo.TransactionLevel,
		TradingKnowledgeLevel:  e.FinancialInfo.TradingKnowledgeLevel,
		CompanyPurpose:         e.FinancialInfo.CompanyPurpose,
		ReferenceRateCompany:   e.FinancialInfo.ReferenceRateCompany,
		RateDate:               e.FinancialInfo.RateDate,
		Rate:                   e.FinancialInfo.Rate,
		CustomerId:             customer.ID,
	})

	for _, item := range e.SejamAddresses {

		_, err = models.CustomerAddress{}.GetOrCreate(models.CustomerAddress{
			CountryId:                 item.CountryId,
			CountryName:               item.CountryName,
			ProvinceName:              item.ProvinceName,
			ProvinceId:                item.ProvinceId,
			CityId:                    item.CityId,
			CityName:                  item.CityName,
			SectionId:                 item.SectionId,
			SectionName:               item.SectionName,
			CityPrefix:                item.CityPrefix,
			Address:                   item.RemnantAddress,
			Alley:                     item.Alley,
			Plaque:                    item.Plaque,
			Tel:                       item.Tel,
			CountryPrefix:             item.CountryPrefix,
			Mobile:                    item.Mobile,
			EmergencyTel:              item.EmergencyTel,
			EmergencyTelCityPrefix:    item.EmergencyTelCityPrefix,
			EmergencyTelCountryPrefix: item.EmergencyTelCountryPrefix,
			FaxPrefix:                 item.FaxPrefix,
			Fax:                       item.Fax,
			Website:                   item.Website,
			Email:                     item.Email,
			PostalCode:                item.PostalCode,
			MebbcoType:                "broker",
			CustomerId:                customer.ID,
		})
	}
	for _, item := range e.SejamBankAccounts {
		isDefault := 0
		if item.IsDefault {
			isDefault = 1
		}
		_, err = models.BankAccount{}.GetOrCreate(models.BankAccountParam{
			AccountNumber:  item.AccountNumber,
			BaTypeName:     item.AccountType,
			Shaba:          item.Sheba,
			IsDefault:      int64(isDefault),
			BankId:         int64(item.BankId),
			BankName:       item.BankName,
			BranchName:     item.BranchName,
			BankBranchCode: item.BranchCode,
			CustomerId:     customer.ID,
		})
	}

	return &result, nil
}
