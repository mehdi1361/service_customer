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
			Id:         int32(v.ID),
			Type:       v.Type,
			FirstPart:  v.FirstPart,
			SecondPart: v.SecondPart,
			Code:       v.Code,
			IsDefault:  v.IsDefault,
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
		Id:         int32(tradingCode.ID),
		Type:       tradingCode.Type,
		FirstPart:  tradingCode.FirstPart,
		SecondPart: tradingCode.SecondPart,
		Code:       tradingCode.Code,
		IsDefault:  tradingCode.IsDefault,
	}
	return &result, nil
}

func (Server) CreateTradingCode(ctx context.Context, e *service.CreateTradingCodeItemRequest) (*service.TradingCodeItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	tradingCode := models.TradingCode{
		Type:       e.Type,
		FirstPart:  e.FirstPart,
		SecondPart: e.SecondPart,
		Code:       e.Code,
		IsDefault:  e.IsDefault,
	}
	db.Create(&tradingCode)
	result := service.TradingCodeItemResponse{
		Id:         int32(tradingCode.ID),
		Type:       tradingCode.Type,
		FirstPart:  tradingCode.FirstPart,
		SecondPart: tradingCode.SecondPart,
		Code:       tradingCode.Code,
		IsDefault:  tradingCode.IsDefault,
	}
	return &result, nil
}

func (Server) DeleteTradingCode(ctx context.Context, e *service.TradingCodeItemRequest) (*service.DeleteTradingCodeItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.TradingCode{}, e.Id)

	result := service.DeleteTradingCodeItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
