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
	SejamReferenceCode string          `json:"sejam_reference_code" gorm:"size:100"`
	NormalNationalCode string          `json:"normal_national_code" gorm:"size:11;unique"`
	UserName           string          `json:"user_name" gorm:"size:100"`
	Password           string          `json:"user_name" gorm:"size:100"`
	IsActive           bool            `json:"is_active"`
	IsRayanService     bool            `json:"is_rayan_service"`
	PhonePerson        []PhonePerson   `json:"phone_persons"`
	PrivateInfo        CustomerPrivate `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	LegalInfo          CustomerLegal   `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	JobInfos           []JobInfo       `gorm:"foreignKey:CustomerId"`
	Fund               Fund            `gorm:"foreignKey:CustomerServiceId"`
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

func (c Customer) SetFund(db *gorm.DB, customer service.FundCustomerList, wg *sync.WaitGroup, fundName string,) error {
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
