package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) ProvinceAll(ctx context.Context, e *empty.Empty) (*service.ProvinceListResponse, error) {

	var data []*service.ProvinceItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allProvince := []models.Province{}
	db.Find(&allProvince)

	for _, v := range allProvince {
		data = append(
			data, &service.ProvinceItemResponse{
				Id:   int32(v.ID),
				CountryId: int32(v.CountryId),
				ProvinceId: int32(v.ProvinceId),
				Name: v.Name,
			},
		)
	}

	result := service.ProvinceListResponse{
		Provinceitemresponse: data,
	}

	return &result, nil
}


func (Server) GetProvinceById(ctx context.Context, e *service.ProvinceItemRequest) (*service.ProvinceItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	province := models.Province{}
	db.Find(&province, "id=?", e.Id)

	result := service.ProvinceItemResponse{
		Id:   int32(province.ID),
		CountryId:   int32(province.CountryId),
		Name: province.Name,
	}
	return &result, nil
}

func (Server) CreateProvince(ctx context.Context, e *service.CreateProvinceItemRequest) (*service.ProvinceItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	province := models.Province{
		Name: e.Name,
		CountryId: uint(e.Countryid),
		ProvinceId: uint(e.Provinceid),
	}
	db.Create(&province)
	result := service.ProvinceItemResponse{
		Id:   int32(province.ID),
		CountryId:   int32(province.CountryId),
		ProvinceId:   int32(province.ProvinceId),
		Name: province.Name,
	}
	return &result, nil
}

func (Server) DeleteProvince(ctx context.Context, e *service.ProvinceItemRequest) (*service.DeleteProvinceItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Province{}, e.Id)

	result := service.DeleteProvinceItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
