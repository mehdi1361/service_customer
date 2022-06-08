package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerGroupAll(ctx context.Context, e *empty.Empty) (*service.CustomerGroupListResponse, error) {

	var data []*service.CustomerGroupItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allDomain := []models.CustomerGroup{}
	db.Find(&allDomain)

	for _, v := range allDomain {
		data = append(data, &service.CustomerGroupItemResponse{
			Id: int32(v.ID),
			IdRayan: int32(v.RayanId),
			Title: v.Title,
		},
		)
	}

	result := service.CustomerGroupListResponse{
		Customergroupitemresponse: data,
	}

	return &result, nil
}

func (Server) GetCustomerGroupById(ctx context.Context, e *service.CustomerGroupItemRequest) (*service.CustomerGroupItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	domain := models.Domain{}
	db.Find(&domain, "id=?", e.Id)
	result := service.CustomerGroupItemResponse{
		Id:   int32(domain.ID),
		IdRayan:   int32(domain.RayanId),
		Title: domain.Title,
	}
	return &result, nil
}

func (Server) CreateCustomerGroup(ctx context.Context, e *service.CreateCustomerGroupItemRequest) (*service.CustomerGroupItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	domain:= models.CustomerGroup{
		Title: e.Title,
		RayanId: int64(e.IdRayan),
	}
	db.Create(&domain)
	result := service.CustomerGroupItemResponse{
		Id:   int32(domain.ID),
		IdRayan:   int32(domain.RayanId),
		Title: domain.Title,
	}
	return &result, nil
}

func (Server) DeleteCustomerGroup(ctx context.Context, e *service.CustomerGroupItemRequest) (*service.DeleteCustomerGroupItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.CustomerGroup{}, e.Id)

	result := service.DeleteCustomerGroupItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
