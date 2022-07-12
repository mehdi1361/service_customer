package server

import (
	"fmt"
	//	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) SejamRegisterPrivatePerson(ctx context.Context, e *service.SejamRegisterPrivatePersonRequest) (*service.SejamRegisterPrivatePersonResponse, error) {

	result := service.SejamRegisterPrivatePersonResponse{
		Id:      1,
		Message: "hello world",
	}

	isSejami := false

	if e.Profile.Status == "Sejami" {
		isSejami = true
	}

	_, err := models.Customer{}.GetOrCreate(models.Customer{
		NormalNationalCode: e.Profile.NationalCode,
		IsSejami:           isSejami,
		SejamType:          e.Profile.CustomerSejamType,
	})

	if err != nil {
		return &service.SejamRegisterPrivatePersonResponse{
			Id:      400,
			Message: fmt.Sprintf("error: %s", err),
		}, nil
	}

	_, err = models.JobInfo{}.GetOrCreate(models.JobInfo{
		EmploymentDate: e.JobInfo.EmploymentDate,
		CompanyName: e.JobInfo.CompanyName,
		CompanyAddress: e.JobInfo.CompanyAddress,
		CompanyPostalCode: e.JobInfo.CompanyPostalCode,
		CompanyEmail: e.JobInfo.CompanyEmail,
		CompanyWebSite: e.JobInfo.CompanyWebSite,
		CompanyCityPrefix: e.JobInfo.CompanyCityPrefix,
		CompanyPhone: e.JobInfo.CompanyPhone,
		Position: e.JobInfo.Position,
		CompanyFaxPrefix: e.JobInfo.CompanyFaxPrefix,
		CompanyFax: e.JobInfo.CompanyFax,
		JobId: uint(e.JobInfo.JobId),
		JobTitle: e.JobInfo.JobTitle,
		JobDescription: e.JobInfo.JobDescription,

	})


	return &result, nil
}
