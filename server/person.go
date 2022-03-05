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
	db.Find(&comexVisitor, "id=?", e.Id)
	result := service.ComexVisitorItemResponse{
		Id:           int32(comexVisitor.ID),
		ComexIdRayan: int32(comexVisitor.ComexIdRayan),
		FullName:     comexVisitor.FullName,
		TypeMebbcoId: int32(comexVisitor.TypeMebbcoId),
	}
	return &result, nil
}

func (Server) CreateComexVisitor(ctx context.Context, e *service.CreateComexVisitorItemRequest) (*service.ComexVisitorItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	comexVisitor := models.ComexVisitor{
		ComexIdRayan: e.ComexIdRayan,
		FullName:     e.FullName,
		Rate:         e.Rate,
		TypeMebbcoId: uint(e.TypeMebbcoId),
	}
	db.Create(&comexVisitor)
	result := service.ComexVisitorItemResponse{
		Id:           int32(comexVisitor.ID),
		ComexIdRayan: int32(comexVisitor.ComexIdRayan),
		FullName:     comexVisitor.FullName,
		TypeMebbcoId: int32(comexVisitor.TypeMebbcoId),
	}
	return &result, nil
}

func (Server) DeleteComexVisitor(ctx context.Context, e *service.ComexVisitorItemRequest) (*service.DeleteComexVisitorItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.ComexVisitor{}, e.Id)

	result := service.DeleteComexVisitorItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
