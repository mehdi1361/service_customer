package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)
func (Server) CountryAll(ctx context.Context, e *empty.Empty) (*service.CountryListResponse, error) {

	var data []*service.CountryItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allGender := []models.Country{}
	db.Find(&allGender)

	for _, v := range allGender {
		data = append(
			data, &service.CountryItemResponse{
				Id:   int32(v.ID),
				Countryid: int32(v.CountryId),
				Name: v.Name,
			},
		)
	}

	result := service.CountryListResponse{
		Countryitemresponse: data,
	}

	return &result, nil
}

func (Server) GetCountryById(ctx context.Context, e *service.CountryItemRequest) (*service.CountryItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	country := models.Country{}
	db.Find(&country, "id=?", e.Id)

	result := service.CountryItemResponse{
		Id:   int32(country.ID),
		Countryid:   int32(country.CountryId),
		Name: country.Name,
	}
	return &result, nil
}

func (Server) CreateCountry(ctx context.Context, e *service.CreateCountryItemRequest) (*service.CountryItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	country := models.Country{Name: e.Name, CountryId: uint(e.Countryid)}
	db.Create(&country)
	result := service.CountryItemResponse{
		Id:   int32(country.ID),
		Countryid:   int32(country.CountryId),
		Name: country.Name,
	}
	return &result, nil
}

func (Server) DeleteCountry(ctx context.Context, e *service.CountryItemRequest) (*service.DeleteCountryItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Country{}, e.Id)

	result := service.DeleteCountryItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
