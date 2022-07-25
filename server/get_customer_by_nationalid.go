package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) GetPersonByNationalId(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.PersonByNationalIdResponse, error) {


	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customer := models.Customer{}
	person := models.CustomerPrivate{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&person, "customer_id=?", customer.ID)

	return &service.PersonByNationalIdResponse{
		FirstName: person.FirstName,
		LastName: person.LastName,
		FatherName: person.FatherName,
		SeriShChar: person.SeriShChar,
		SeriSh: person.SeriSh,
		Serial: person.Serial,
		ShNumber: person.ShNumber,
		BirthDate: person.BirthDate,
		PlaceOfIssue: person.PlaceOfIssue,
		PlaceOfBirth: person.PlaceOfBirth,
		EconomicCode: "1",
		NationalId: customer.NormalNationalCode,
	}, nil
}
