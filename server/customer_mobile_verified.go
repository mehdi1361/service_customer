package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
	"time"
)

func (Server) CustomerVerified(ctx context.Context, e *service.CustomerVerifiedRequest) (*service.StateResponse, error) {
	now := time.Now()
	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	verificationCode := models.VerificationCode{}

	if err := db.Where("customer_id=? and is_active=?", customer.ID, true).First(&verificationCode).Error; err != nil {

		return &service.StateResponse{
			Id:      404,
			Message: "verification code not found",
		}, nil
	}

	if int(now.Sub(verificationCode.CreatedAt).Minutes()) > 2 {

		return &service.StateResponse{
			Id:      400,
			Message: "verification code expired",
		}, nil
	}

	if e.VerificationCode != verificationCode.Code {
		return &service.StateResponse{
			Id:      400,
			Message: "verification code not match!!!",
		}, nil

	}

	db.Model(&models.VerificationCode{}).Where(
		"customer_id=? and is_active=? and code=?", customer.ID, true,e.VerificationCode,
	).Update("is_active", false)

	return &service.StateResponse{
		Id:      200,
		Message: "verification code accepted",
	}, nil

}
