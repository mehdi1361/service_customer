package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
	"strconv"
)

func (Server) CustomerGetActiveMobile(ctx context.Context, e *service.CustomerActiveMobileRequest) (*service.CustomerActiveMobileResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	phones := []*models.PhonePerson{}

	db.Find(&phones, "customer_id=? and is_mobile=?", customer.ID, true)

	for _, v := range phones {
		if (v.IsActive) {
			return &service.CustomerActiveMobileResponse{
				Customermobile: &service.CustomerMobile{
					PhoneNumber: v.PhoneNumber,
					PhoneId:     strconv.Itoa(int(v.ID)),
					IsActive:    v.IsActive,
				},
			}, nil

		}

	}
	return nil, nil

}
