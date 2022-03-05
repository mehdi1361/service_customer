package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) ComexVisitorAll(ctx context.Context, e *empty.Empty) (*service.ComexVisitorListResponse, error) {

	var data []*service.ComexVisitorItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	comexVisitorAll := []models.ComexVisitor{}
	db.Find(&comexVisitorAll)

	for _, v := range comexVisitorAll {
		data = append(data, &service.ComexVisitorItemResponse{
			Id: int32(v.ID),
			ComexIdRayan: int32(v.ComexIdRayan),
			FullName: v.FullName,
			TypeMebbcoId: int32(v.TypeMebbcoId),
		},
		)
	}

	result := service.ComexVisitorListResponse{
		Comexvisitoritemresponse: data,
	}

	return &result, nil
}

func (Server) GetComexVisitorById(ctx context.Context, e *service.ComexVisitorItemRequest) (*service.ComexVisitorItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	comexVisitor := models.ComexVisitor{}
	db.Find(&comexVisitor, "id=?", e.Id)
	result := service.ComexVisitorItemResponse{
			Id: int32(comexVisitor.ID),
			ComexIdRayan: int32(comexVisitor.ComexIdRayan),
			FullName: comexVisitor.FullName,
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
	comexVisitor:= models.ComexVisitor{
		ComexIdRayan: e.ComexIdRayan,
		FullName: e.FullName,
		Rate: e.Rate,
		TypeMebbcoId: uint(e.TypeMebbcoId),
	}
	db.Create(&comexVisitor)
	result := service.ComexVisitorItemResponse{
			Id: int32(comexVisitor.ID),
			ComexIdRayan: int32(comexVisitor.ComexIdRayan),
			FullName: comexVisitor.FullName,
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
