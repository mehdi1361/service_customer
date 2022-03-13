package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//"github.com/k0kubun/pp/v3"
	service "service_customer/service/rayan/proto"
	"strings"
	"sync"
	"time"
)

type Customer struct {
	gorm.Model
	SejamReferenceCode string          `json:"sejam_reference_code" gorm:"size:100"`
	NormalNationalCode string          `json:"normal_national_code gorm:"size:10"`
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

func (c Customer) SetBulkDataFund(customers []*service.FundCustomerList) {
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
		c.SetFund(db, *v, &wg)
	}
	fmt.Printf("length of customer lis is %d", len(customers))
	wg.Wait()
	defer db.Close()
}

func (c Customer) SetFund(
	db *gorm.DB,
	customer service.FundCustomerList,
	wg *sync.WaitGroup,
) error {
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
	newCustomer := Customer{
		NormalNationalCode: normalNationalCode,
		IsActive:           true,
		IsRayanService:     true,
	}

	if err := tx.Create(&newCustomer).Error; err != nil {
		tx.Rollback()
		return err
	}

	fund := Fund{
		CustomerServiceId:  newCustomer.ID,
		CustomerId:         customer.CustomerId,
		ReferredBy:         customer.ReferredBy,
		Personality:        customer.Personality,
		FullName:           customer.CustomerFullName,
		IssunigCity:        customer.IssuingCity,
		Nationality:        customer.Nationality,
		NationalIdentifier: customer.NationalIdentifier,
	}

	if err := tx.Create(&fund).Error; err != nil {
		tx.Rollback()
		return err
	}

	//pp.Print(customer)
	return tx.Commit().Error
}
