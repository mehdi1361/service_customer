package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) BankAll(ctx context.Context, e *empty.Empty) (*service.BankListResponse, error) {

	var data []*service.BankItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allBank := []models.Bank{}
	db.Find(&allBank)

	for _, v := range allBank {
		data = append(data, &service.BankItemResponse{
			Id: int32(v.ID),
			BankId: int32(v.BankId),
			Title: v.Title,
		},
		)
	}

	result := service.BankListResponse{
		Bankitemresponse: data,
	}

	return &result, nil
}

func (Server) GetBankById(ctx context.Context, e *service.BankItemRequest) (*service.BankItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	bank := models.Bank{}
	db.Find(&bank, "id=?", e.Id)
	result := service.BankItemResponse{
		Id:   int32(bank.ID),
		BankId:   int32(bank.BankId),
		Title: bank.Title,
	}
	return &result, nil
}

func (Server) CreateBank(ctx context.Context, e *service.CreateBankItemRequest) (*service.BankItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	bank:= models.Bank{
		Title: e.Title,
		BankId: uint(e.BankId),
	}
	db.Create(&bank)
	result := service.BankItemResponse{
		Id:   int32(bank.ID),
		Title: bank.Title,
	}
	return &result, nil
}

func (Server) DeleteBank(ctx context.Context, e *service.BankItemRequest) (*service.DeleteBankItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Bank{}, e.Id)

	result := service.DeleteBankItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
