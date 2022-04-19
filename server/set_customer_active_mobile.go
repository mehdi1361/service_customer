package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerSetActiveMobile(ctx context.Context, e *service.CustomerSetActiveMobileRequest) (*service.StateResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)


	db.Model(&models.PhonePerson{}).Where("customer_id=? and is_mobile=?", customer.ID, true).Update("is_active", false)

	db.Model(&models.PhonePerson{}).Where(
		"customer_id=? and is_mobile=? and phone_number=?", customer.ID, true,e.PhoneNumber,
	).Update("is_active", true)

	return &service.StateResponse{
		Id: 200,
		Message: "activation mobile set",
	}, nil

}
