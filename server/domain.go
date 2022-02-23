package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) DomainAll(ctx context.Context, e *empty.Empty) (*service.DomainListResponse, error) {

	var data []*service.DomainItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allDomain := []models.Domain{}
	db.Find(&allDomain)

	for _, v := range allDomain {
		data = append(data, &service.DomainItemResponse{
			Id: int32(v.ID),
			IdRayan: int32(v.IdRayan),
			Title: v.Title,
		},
		)
	}

	result := service.DomainListResponse{
		Domainitemresponse: data,
	}

	return &result, nil
}

func (Server) GetDomainById(ctx context.Context, e *service.DomainItemRequest) (*service.DomainItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	domain := models.Domain{}
	db.Find(&domain, "id=?", e.Id)
	result := service.DomainItemResponse{
		Id:   int32(domain.ID),
		IdRayan:   int32(domain.IdRayan),
		Title: domain.Title,
	}
	return &result, nil
}

func (Server) CreateDomain(ctx context.Context, e *service.CreateDomainItemRequest) (*service.DomainItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	domain:= models.Domain{
		Title: e.Title,
		IdRayan: e.IdRayan,
	}
	db.Create(&domain)
	result := service.DomainItemResponse{
		Id:   int32(domain.ID),
		IdRayan:   domain.IdRayan,
		Title: domain.Title,
	}
	return &result, nil
}

func (Server) DeleteDomian(ctx context.Context, e *service.DomainItemRequest) (*service.DeleteDomainItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Domain{}, e.Id)

	result := service.DeleteDomainItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
