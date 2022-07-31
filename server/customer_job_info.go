package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) GetPersonJobInfo(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.JobInfoResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	customer := models.Customer{}
	jobInfo := models.JobInfo{}
	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&jobInfo, "customer_id=?", customer.ID)

	return &service.JobInfoResponse{
		EmploymentDate:    jobInfo.EmploymentDate,
		CompanyName:       jobInfo.CompanyName,
		CompanyAddress:    jobInfo.CompanyAddress,
		CompanyPostalCode: jobInfo.CompanyPostalCode,
		CompanyEmail:      jobInfo.CompanyEmail,
		CompanyWebSite:    jobInfo.CompanyWebSite,
		CompanyCityPrefix: jobInfo.CompanyCityPrefix,
		CompanyPhone:      jobInfo.CompanyPhone,
		Position:          jobInfo.Position,
		CompanyFaxPrefix:  jobInfo.CompanyFaxPrefix,
		CompanyFax:        jobInfo.CompanyFax,
		JobTitle:          jobInfo.JobTitle,
		JobDescription:    jobInfo.JobDescription,
	}, nil
}
