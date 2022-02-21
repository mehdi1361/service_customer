package main

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func main() {
	grpcServer := grpc.NewServer()
	var server Server
	service.RegisterCustomerServiceServer(grpcServer, server)
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
} // Server is implementation proto interface

type Server struct{} // Search function responsible to get the Country information
func (Server) GenderAll(ctx context.Context, e *empty.Empty) (*service.GenderListResponse, error) {

	var data []*service.GenderItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	allGender := []models.Gender{}
	db.Find(&allGender)

	for _, v := range allGender {
		data = append(data, &service.GenderItemResponse{Id: int32(v.ID), Name: v.Name})
	}

	result := service.GenderListResponse{
		Genderitemresponse: data,
	}

	return &result, nil
}

func (Server) GetGenderById(ctx context.Context, e *service.GenderItemRequest) (*service.GenderItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	gender := models.Gender{}
	db.Find(&gender, "id=?", e.Id)
	result := service.GenderItemResponse{
		Id:   int32(gender.ID),
		Name: gender.Name,
	}
	return &result, nil
}

func (Server) CreateGender(ctx context.Context, e *service.CreateGenderItemRequest) (*service.GenderItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	gender := models.Gender{Name: e.Name}
	db.Create(&gender)
	result := service.GenderItemResponse{
		Id:   int32(gender.ID),
		Name: gender.Name,
	}
	return &result, nil
}

func (Server) DeleteGender(ctx context.Context, e *service.GenderItemRequest) (*service.DeleteGenderItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Gender{}, e.Id)

	result := service.DeleteGenderItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}

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
