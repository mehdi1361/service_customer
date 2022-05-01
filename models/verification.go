package models

import (
	"github.com/jinzhu/gorm"
	//"golang.org/x/net/context"
	//sms "service_customer/service/sms/proto"
	//service "service_customer/service_connections"
	//	utils "service_customer/utils"
)

type VerificationCode struct {
	gorm.Model
	Code       string `json:"name" gorm:"size:4"`
	IsActive   bool   `json:"is_active" gorm:"Column:is_active;default:true"`
	CustomerId uint   `json:"customer_id"`
}

func (v *VerificationCode) TableName() string {
	return "customer_verification"
}

func (v VerificationCode) SendVerificationCode(customer Customer) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Model(&VerificationCode{}).Where("customer_id=?", customer.ID).Update("is_active", false)
	//verificationCode := utils.RandomCodeGenerate(4)
	verificationCode := "1000"

	phone := PhonePerson{}
	db.Find(&phone, "customer_id=? and is_active=?", customer.ID, true)

	/*	conn, _ := service.SmsServiceConnection()
	ce := sms.NewSmsServiceClient(conn)
	_, err = ce.SmsService(
		context.Background(),
		&sms.SmsSendRequest{
			NameService: "customer service",
			Phone:       []string{phone.PhoneNumber},
			Text:        []string{fmt.Sprintf("کد فعالسازی سامانه کارگزاری بانک خاورمیانه %s", verificationCode)},
		},
	)
	*/

	verificarionCode := VerificationCode{
		Code:       verificationCode,
		CustomerId: customer.ID,
		IsActive:   true,
	}
	db.Create(&verificarionCode)

	return nil
}
