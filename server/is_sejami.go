package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerIsSejami(ctx context.Context, e *service.CustomerIsSejamiRequest) (*service.CustomerSejamiResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=? and is_sejami=?", e.NationalId, e.IsSejami)

	if customer.ID != 0 {
		return &service.CustomerSejamiResponse{
			Exist: true,
		}, nil
	}
	return &service.CustomerSejamiResponse{
		Exist: false,
	}, nil

}
