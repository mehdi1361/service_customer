package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	service "service_customer/service/rayan/proto"
	"strings"
	"sync"
	"time"
)

type Customer struct {
	gorm.Model
	SejamReferenceCode string             `json:"sejam_reference_code" gorm:"size:100"`
	NormalNationalCode string             `json:"normal_national_code" gorm:"size:11;unique"`
	UserName           string             `json:"user_name" gorm:"size:100"`
	Password           string             `json:"user_name" gorm:"size:100"`
	IsActive           bool               `json:"is_active"`
	IsRayanService     bool               `json:"is_rayan_service"`
	PhonePerson        []PhonePerson      `json:"phone_persons"`
	PrivateInfo        CustomerPrivate    `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	LegalInfo          CustomerLegal      `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	JobInfos           []JobInfo          `gorm:"foreignKey:CustomerId"`
	VerificationCodes  []VerificationCode `gorm:"foreignKey:CustomerId"`
	Fund               Fund               `gorm:"foreignKey:CustomerServiceId"`
	IsSejami           bool               `json:"is_sejami" gorm:"Column:is_sejami;"`
	SejamType          string             `json:"sejam_type gorm:"size:100;Column:sejam_type"`
}

func (c *Customer) TableName() string {
	return "customer_customer"
}

func (c Customer) SetBulkDataFund(customers []*service.FundCustomerList, fundName string) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	db, err := Connect()
	sqlDB := db.DB()
	if err != nil {
		fmt.Print(err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(300)

	sqlDB.SetConnMaxLifetime(time.Hour)

	for _, v := range customers {

		wg.Add(1)
		c.SetFund(db, *v, &wg, fundName)
	}
	wg.Wait()
	defer db.Close()
}

func (c Customer) SetFund(db *gorm.DB, customer service.FundCustomerList, wg *sync.WaitGroup, fundName string) error {
	defer wg.Done()
	sqlDB := db.DB()
	for {
		if e := sqlDB.Ping(); e == nil {
			break
		}
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()

		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	normalNationalCode := strings.Replace(customer.NationalIdentifier, "-", "", -1)
	fCustomer := Customer{}

	db.Find(&fCustomer, "normal_national_code=?", normalNationalCode)

	if fCustomer.ID == 0 {
		fCustomer = Customer{
			NormalNationalCode: normalNationalCode,
			IsActive:           true,
			IsRayanService:     true,
		}

		if err := tx.Create(&fCustomer).Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	fFund := Fund{}
	db.Find(&fFund, "customer_service_id=? and fund_name=?", fCustomer.ID, fundName)

	if fFund.ID == 0 {
		fund := Fund{
			CustomerServiceId:  fCustomer.ID,
			CustomerId:         customer.CustomerId,
			ReferredBy:         customer.ReferredBy,
			Personality:        customer.Personality,
			FullName:           customer.CustomerFullName,
			IssunigCity:        customer.IssuingCity,
			Nationality:        customer.Nationality,
			NationalIdentifier: customer.NationalIdentifier,
			FundName:           fundName,
		}

		if err := tx.Create(&fund).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	//pp.Print(customer)
	return tx.Commit().Error
}

func (c Customer) SetBulkDataBroker(customers []*service.BrokerCustomerList) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	db, err := Connect()
	sqlDB := db.DB()
	if err != nil {
		fmt.Print(err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(300)

	sqlDB.SetConnMaxLifetime(time.Hour)

	for _, v := range customers {

		wg.Add(1)
		c.SetBroker(db, *v, &wg)
	}
	wg.Wait()
	defer db.Close()
}

func (c Customer) SetBroker(db *gorm.DB, customer service.BrokerCustomerList, wg *sync.WaitGroup) error {
	defer wg.Done()
	sqlDB := db.DB()
	for {
		if e := sqlDB.Ping(); e == nil {
			break
		}
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()

		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	normalNationalCode := strings.Replace(customer.NationalCode, "-", "", -1)
	fCustomer := Customer{}

	db.Find(&fCustomer, "normal_national_code=?", normalNationalCode)

	if fCustomer.ID == 0 {
		fCustomer = Customer{
			NormalNationalCode: normalNationalCode,
			IsActive:           true,
			IsRayanService:     true,
		}

		if err := tx.Create(&fCustomer).Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	brokerData := Broker{}
	db.Find(&brokerData, "customer_id=?", fCustomer.ID)

	if brokerData.ID == 0 {
		broker := Broker{
			CustomerServiceId:          fCustomer.ID,
			CustomerId:                 customer.CustomerId,
			TelegramUsername:           customer.TelegramUsername,
			TelegramStatusId:           customer.TelegramStatusId,
			BourseAccountName:          customer.BourseAccountName,
			AccountNumber:              customer.AccountNumber,
			OnlineUsername:             customer.OnlineUsername,
			HasOnlineAccount:           customer.HasOnlineAccount,
			ModificationDate:           customer.ModificationDate,
			IsMmtpUser:                 customer.IsMmtpUser,
			MmtpUserStatusId:           customer.MmtpUserStatusId,
			IsSiteUser:                 customer.IsSiteUser,
			EorderStatusId:             customer.EorderStatusId,
			HasSignSample:              customer.HasSignSample,
			HasCustomerPhoto:           customer.HasCustomerPhoto,
			HasBirthCertificate:        customer.HasBirthCertificate,
			HasCertificateComments:     customer.HasCertificateComments,
			HasZipFile:                 customer.HasZipFile,
			HasOfficialGazette:         customer.HasOfficialGazette,
			HasOfficialAds:             customer.HasOfficialAds,
			ComexVisitorId:             customer.ComexVisitorId,
			MmtpUserId:                 customer.MmtpUserId,
			ComexEconomyAccount:        customer.ComexEconomyAccount,
			IsPortfo:                   customer.IsPortfo,
			TraderCredit:               customer.TraderCredit,
			ComexCredit:                customer.ComexCredit,
			SfCredit:                   customer.SfCredit,
			Credit:                     customer.Credit,
			IsStockCreditPurchase:      customer.IsStockCreditPurchase,
			IsCollateralStocksCustomer: customer.IsCollateralStocksCustomer,
			CustomerIdentity:           customer.CustomerIdentity,
		}

		if err := tx.Create(&broker).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (c Customer) GetOrCreate(d Customer) (*Customer, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	normalNationalCode := strings.Replace(d.NormalNationalCode, "-", "", -1)
	fCustomer := Customer{}

	db.Find(&fCustomer, "normal_national_code=?", normalNationalCode)

	if fCustomer.ID == 0 {
		fCustomer = Customer{
			NormalNationalCode: normalNationalCode,
			IsSejami:           d.IsSejami,
			SejamType:          d.SejamType,
			IsActive:            true,
		}

		db.Create(&fCustomer)

	}

	return &fCustomer, nil
}
