package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func userConfirmState(customerId, stateId uint) bool {
	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
		return false
	}
	customerState := models.CustomerState{}
	db.Find(&customerState, "customer_id=? and state_id=?", customerId, stateId)

	if customerState.ID != 0 {
		return customerState.Confirm
	}

	defer db.Close()

	return false
}

func (Server) CustomerGetState(ctx context.Context, e *service.CustomerGetStateRequest) (*service.CustomerGetStateResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	customerStates := []*service.CustomerState{}
	customer := models.Customer{}
	states := []models.State{}

	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&states)

	for _, v := range states {
		customerStates = append(customerStates, &service.CustomerState{
			Id:        uint32(v.ID),
			StateName: v.StateName,
			Title:     v.Title,
			IconClass: v.IconClass,
			StateId:   v.StateId,
		})
	}

	return &service.CustomerGetStateResponse{
		CustomerState: customerStates,
	}, nil
}
