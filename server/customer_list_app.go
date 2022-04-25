package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerListApp(ctx context.Context, e *service.CustomerListAppRequest) (*service.CustomerListAppResponse, error) {
	funds := []models.Fund{}
	db, _ := models.Connect()
	inlineResult := []*service.CustomerApp{}
	defer db.Close()


	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	db.Where("customer_service_id=?", customer.ID).Find(&funds)

	for _, v := range funds {
		inlineResult = append(
			inlineResult,
			&service.CustomerApp{
				Id:   v.CustomerId,
				Name: v.FundName,
			})
	}
	return &service.CustomerListAppResponse{
		Customerapp: inlineResult,
	}, nil

}
