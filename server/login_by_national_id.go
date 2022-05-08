package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
	message "service_customer/utils/message"

)

func (Server) LoginByNationalId(ctx context.Context, e *service.LoginByNationalIdRequest) (*service.LoginStateResponse, error) {

	db, err := models.Connect()
	if err != nil {
		return &service.LoginStateResponse{
			Id:      400,
			Message: fmt.Sprintf("error in connect db %s", err),
		}, nil
	}
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	if customer.ID == 0 {

		return &service.LoginStateResponse{
			Id:      404,
			Message: message.UserNotFound,
		}, nil
	}
	var count int64
	db.Model(&models.PhonePerson{}).Where("customer_id=? and is_mobile=? and is_active=?", customer.ID, true, true).Count(&count)

	if count > 1 {
		return &service.LoginStateResponse{
			Id:      403,
			Message: message.MultipleMoobileActive,
		}, nil
	}


	err = models.VerificationCode{}.SendVerificationCode(customer)
	if err != nil {
		return &service.LoginStateResponse{
			Id:      400,
			Message: message.ErrorSendVerificationCode,
		}, nil

	}



	return &service.LoginStateResponse{
		Id:      200,
		Message: message.VerificationCodeSent,
	}, nil
}
