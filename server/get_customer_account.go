package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) GetPersonBankAccount(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.PersonBankAccountResponse, error) {


	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customer := models.Customer{}
	 accounts := []models.BankAccount{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&accounts, "customer_id=?", customer.ID)

	serviceAccounts := []*service.PersonBankAccount{}

	for _, v := range accounts {
		serviceAccounts = append(serviceAccounts, &service.PersonBankAccount{
			AccountNumber: v.AccountNumber,
			BaTypeName: v.BaTypeName,
			Shaba: v.Shaba,
			IsDefault: v.IsDefault,
			IsActive: v.IsActive,
			IsOnline: v.IsOnline,
			BranchId: uint32(v.BranchId),
		})
	}

	return &service.PersonBankAccountResponse{
		PersonBankAccount: serviceAccounts,
	}, nil
}
