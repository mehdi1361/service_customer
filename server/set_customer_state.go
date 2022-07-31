package server

import (
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerUpdateState(ctx context.Context, e *service.CustomerUpdateStateRequest) (*service.CustomerUpdateStateResponse, error) {

	db, _ := models.Connect()
	defer db.Close()
	customer := models.Customer{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)

	if customer.ID == 0 {
		return &service.CustomerUpdateStateResponse{
			Message: "کاربر یافت نشد",
			Code:    404,
		}, nil
	}

	state := models.State{}
	db.Find(&state, "id=?", e.StateId)

	if state.ID == 0 {
		return &service.CustomerUpdateStateResponse{
			Message: "مرحله یافت نشد",
			Code:    404,
		}, nil
	}

	customerState := models.CustomerState{}
	db.Find(&state, "customer_id=? and state_id=?", customer.ID, state.ID)

	if customerState.ID == 0 {
		db.Create(&models.CustomerState{
			CustomerId: customer.ID,
			StateId:    state.ID,
			Confirm:    e.Confirm,
		})
	} else {
		customerState.Confirm = e.Confirm
		db.Save(&customer)
	}

	return &service.CustomerUpdateStateResponse{
		Message: "تایید مرحله",
		Code:    200,
	}, nil

}
