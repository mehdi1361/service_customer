package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) PersonAll(ctx context.Context, e *empty.Empty) (*service.PersonListResponse, error) {

	var data []*service.PersonItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customerAll := []models.Customer{}
	db.Find(&customerAll)

	for _, v := range customerAll {
		data = append(data, &service.PersonItemResponse{
			Id:                 int32(v.ID),
			SejamReferenceCode: v.SejamReferenceCode,
			UserName:           v.UserName,
			Password:           v.Password,
			IsActive:           v.IsActive,
			IsRayanService:     v.IsRayanService,
		},
		)
	}

	result := service.PersonListResponse{
		Personitemresponse: data,
	}

	return &result, nil
}

func (Server) GetPersonById(ctx context.Context, e *service.PersonItemRequest) (*service.PersonItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "id=?", e.Id)
	result := service.PersonItemResponse{
		Id:                 int32(customer.ID),
		SejamReferenceCode: customer.SejamReferenceCode,
		UserName:           customer.UserName,
		Password:           customer.Password,
		IsActive:           customer.IsActive,
		IsRayanService:     customer.IsRayanService,
	}
	return &result, nil
}

func (Server) CreatePerson(ctx context.Context, e *service.CreatePersonItemRequest) (*service.PersonItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	customer := models.Customer{
		SejamReferenceCode: e.SejamReferenceCode,
		UserName:           e.UserName,
		Password:           e.Password,
		IsActive:           e.IsActive,
		IsRayanService:     e.IsRayanService,
	}
	db.Create(&customer)
	result := service.PersonItemResponse{
		Id:                 int32(customer.ID),
		SejamReferenceCode: customer.SejamReferenceCode,
		UserName:           customer.UserName,
		Password:           customer.Password,
		IsActive:           customer.IsActive,
		IsRayanService:     customer.IsRayanService,
	}
	return &result, nil
}

func (Server) DeletePerson(ctx context.Context, e *service.PersonItemRequest) (*service.DeletePersonItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Customer{}, e.Id)

	result := service.DeletePersonItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
