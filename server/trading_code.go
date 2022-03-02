package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) TradingCodeAll(ctx context.Context, e *empty.Empty) (*service.TradingCodeListResponse, error) {

	var data []*service.TradingCodeItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	tradingCodes := []models.TradingCode{}
	db.Find(&tradingCodes)

	for _, v := range tradingCodes {
		data = append(data, &service.TradingCodeItemResponse{
			Id: int32(v.ID),
			Type: v.Type,
			FirstPart: v.FirstPart,
			SecondPart: v.SecondPart,
			Code: v.Code,
			IsDefault: v.IsDefault,
		},
		)
	}

	result := service.TradingCodeListResponse{
		Tradingcodeitemresponse: data,
	}

	return &result, nil
}

func (Server) GetTradingCodeId(ctx context.Context, e *service.TradingCodeItemRequest) (*service.TradingCodeItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	tradingCode := models.TradingCode{}
	db.Find(&tradingCode, "id=?", e.Id)
	result := service.TradingCodeItemResponse{
		Id:   int32(tradingCode.ID),
		Type: tradingCode.Type,
		JobId: job.JobId,
	}
	return &result, nil
}

func (Server) CreateJob(ctx context.Context, e *service.CreateJobItemRequest) (*service.JobItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	job := models.Job{
		Title: e.Title,
		JobId: e.JobId,
	}
	db.Create(&job)
	result := service.JobItemResponse{
		Id:   int32(job.ID),
		Title: job.Title,
		JobId: job.JobId,
	}
	return &result, nil
}

func (Server) DeleteJob(ctx context.Context, e *service.JobItemRequest) (*service.DeleteJobItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Job{}, e.Id)

	result := service.DeleteJobItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
