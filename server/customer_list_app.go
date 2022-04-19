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

	db.Where("customer_service_id=?", e.CustomerId).Find(&funds)

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
