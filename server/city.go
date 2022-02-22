package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CityAll(ctx context.Context, e *empty.Empty) (*service.CityListResponse, error) {

	var data []*service.CityItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allCity := []models.City{}
	db.Find(&allCity)

	for _, v := range allCity {
		data = append(
			data, &service.CityItemResponse{
				Id:   int32(v.ID),
				CityId: int32(v.CityId),
				ProvinceId: int32(v.ProvinceId),
				Name: v.Name,
			},
		)
	}

	result := service.CityListResponse{
		Cityitemresponse: data,
	}

	return &result, nil
}


func (Server) GetCityById(ctx context.Context, e *service.CityItemRequest) (*service.CityItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	city := models.City{}
	db.Find(&city, "id=?", e.Id)

	result := service.CityItemResponse{
		Id:   int32(city.ID),
		CityId:   int32(city.CityId),
		ProvinceId:   int32(city.ProvinceId),
		Name: city.Name,
	}
	return &result, nil
}

func (Server) CreateCity(ctx context.Context, e *service.CreateCityItemRequest) (*service.CityItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	city := models.City{
		Name: e.Name,
		CityId: uint(e.CityId),
		ProvinceId: uint(e.ProvinceId),
	}
	db.Create(&city)
	result := service.CityItemResponse{
		Id:   int32(city.ID),
		CityId:   int32(city.CityId),
		ProvinceId:   int32(city.ProvinceId),
		Name: city.Name,
	}
	return &result, nil
}

func (Server) DeleteCity(ctx context.Context, e *service.CityItemRequest) (*service.DeleteCityItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.City{}, e.Id)

	result := service.DeleteCityItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
