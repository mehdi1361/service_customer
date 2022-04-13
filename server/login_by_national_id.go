package server

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) LoginByNationalId(ctx context.Context, e *service.LoginByNationalIdRequest) (*service.LoginStateResponse, error) {

	db, err := models.Connect()
	if err != nil {
		return &service.LoginStateResponse{
			Id:      400,
			Message: fmt.Sprintf("error in connect db %s", err),
		}, errors.New(fmt.Sprintf("error in connect db %s", err))
	}
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	if customer.ID == 0 {

		return &service.LoginStateResponse{
			Id:      404,
			Message: "user not found",
		}, errors.New("user not found")
	}

	return &service.LoginStateResponse{
		Id:      200,
		Message: "verification code sent",
	}, nil
}
