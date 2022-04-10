package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) MebbcoBranchAll(ctx context.Context, e *empty.Empty) (*service.MebbcoBranchListResponse, error) {

	var data []*service.MebbcoBranchItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	branchs := []models.MebbcoBranch{}
	db.Find(&branchs)

	for _, v := range branchs {
		data = append(data, &service.MebbcoBranchItemResponse{
			Id: int32(v.ID),
			IdRayan: int32(v.IdRayan),
			TypeMebbco: v.TypeMebbco,
			Title: v.Title,
		},
		)
	}

	result := service.MebbcoBranchListResponse{
		Mebbcobranchitemresponse: data,
	}

	return &result, nil
}

func (Server) GetMebbcoBranchById(ctx context.Context, e *service.MebbcoBranchItemRequest) (*service.MebbcoBranchItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	branch := models.MebbcoBranch{}
	db.Find(&branch, "id=?", e.Id)
	result := service.MebbcoBranchItemResponse{
		Id:   int32(branch.ID),
		IdRayan:   int32(branch.IdRayan),
		TypeMebbco: branch.TypeMebbco,
		Title: branch.Title,
	}
	return &result, nil
}

func (Server) CreateMebbcoBranch(ctx context.Context, e *service.CreateMebbcoBranchItemRequest) (*service.MebbcoBranchItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	mebbcoBranch := models.MebbcoBranch{
		Title: e.Title,
		TypeMebbco: e.TypeMebbco,
		IdRayan: int(e.IdRayan),
	}
	db.Create(&mebbcoBranch)
	result := service.MebbcoBranchItemResponse{
		Id:   int32(mebbcoBranch.ID),
		IdRayan:   int32(mebbcoBranch.IdRayan),
		Title: mebbcoBranch.Title,
		TypeMebbco: mebbcoBranch.TypeMebbco,
	}
	return &result, nil
}

func (Server) DeleteMebbcoBranch(ctx context.Context, e *service.MebbcoBranchItemRequest) (*service.DeleteMebbcoBranchItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Domain{}, e.Id)

	result := service.DeleteMebbcoBranchItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
