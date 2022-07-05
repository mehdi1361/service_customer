package server

import (
	//	"fmt"
	//	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	//	models "service_customer/models"
	service "service_customer/service/proto"
)


func (Server) SejamRegisterPrivatePerson(ctx context.Context, e *service.SejamRegisterPrivatePersonRequest) (*service.SejamRegisterPrivatePersonResponse, error) {

	result := service.SejamRegisterPrivatePersonResponse{
		Id:   1,
		Message: "hello world",
	}
	return &result, nil
}
