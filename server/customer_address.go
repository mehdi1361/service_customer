package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) GetPersonByAddress(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.PersonByAddressResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customer := models.Customer{}
	address := models.CustomerAddress{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&address, "customer_id=?", customer.ID)

	return &service.PersonByAddressResponse{
		PostalCode:   address.PostalCode,
		Address:      address.Address,
		Mobile:       address.Mobile,
		Fax:          address.Fax,
		Tel:          address.Tel,
		Email:        address.Email,
		ProvinceName: address.ProvinceName,
		CityName:     address.CityName,
	}, nil
}
