package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
	utils "service_customer/utils"
	"strconv"
)

func (Server) CustomerMobileList(ctx context.Context, e *service.LoginByNationalIdRequest) (*service.CustomerMobileListResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	phones := []*models.PhonePerson{}

	db.Find(&phones, "customer_id=? and is_mobile=?", customer.ID, true)
	fmt.Println(len(phones))

	servicePhones := []*service.CustomerMobile{}
	for _, v := range phones {
		servicePhones = append(servicePhones, &service.CustomerMobile{
			PhoneNumber: utils.HashMobile(v.PhoneNumber),
			PhoneId: strconv.Itoa(int(v.ID)),
			IsActive: v.IsActive,
		})
	}
	return &service.CustomerMobileListResponse{
		Customermobile: servicePhones,
	}, nil

}
