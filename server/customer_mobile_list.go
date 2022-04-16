package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerMobileList(ctx context.Context, e *service.LoginByNationalIdRequest) (*service.CustomerMobileListResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	phones := []*models.PhonePerson{}

	db.Find(&phones, "customer_id=? and is_mobile=? and is_active=?", customer.ID, true, true)
	fmt.Println(len(phones))

	servicePhones := []*service.CustomerMobile{}
	for _, v := range phones {
		servicePhones = append(servicePhones, &service.CustomerMobile{
			PhoneNumber: v.PhoneNumber,
			IsActive: v.IsActive,
		})
	}
	return &service.CustomerMobileListResponse{
		Customermobile: servicePhones,
	}, nil

}
