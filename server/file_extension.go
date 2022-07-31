package server

import (
	"fmt"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) CustomerGetFile(ctx context.Context, e *service.PersonByNationalIdRequest) (*service.CustomerGetFileResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	customerServiceFiles := []*service.CustomerGetFile{}
	customer := models.Customer{}
	fileTypes := []models.FileType{}
	cFile := models.CustomerFile{}

	db.Find(&customer, "normal_national_code=?", e.NationalId)
	db.Find(&fileTypes)

	for _, v := range fileTypes {
		db.Find(&cFile, "customer_id=? and file_type_id=?", customer.ID, v.ID)

		customerServiceFiles = append(customerServiceFiles, &service.CustomerGetFile{
			Id:        uint32(v.ID),
			Name: v.Name,
			FaName: v.FaName,
			IsForce:   v.IsForce,
			FileData: cFile.FileData,
			FileExtension: &service.FileExtension{
				Name: "image/jpg, image/jpeg, image/png",
				RealSize: 5 * 1024 * 1024,
				Size: "5MB",
			},
		})
	}

	return &service.CustomerGetFileResponse{
		CustomerGetFile: customerServiceFiles,
	}, nil
}
