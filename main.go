package main

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	service "service_customer/service/proto"
	models "service_customer/models"
)

func main() {
	grpcServer := grpc.NewServer()
	var server Server
	service.RegisterGenderServer(grpcServer, server)
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
	fmt.Println(allGender)


	for _, v := range allGender {
		data = append(data, &service.GenderItemResponse{Id: int32(v.ID), Name: v.Name})
	}

	result := service.GenderListResponse{
		Genderitemresponse: data,
	}
	fmt.Println(result)


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
		Id: int32(gender.ID),
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
		Id: int32(gender.ID),
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
