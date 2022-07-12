package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CustomerPrivate struct {
	gorm.Model

	FirstName     string `json:"first_name" gorm:"size:60"`
	LastName      string `json:"last_name" gorm:"size:60"`
	FatherName    string `json:"father_name" gorm:"size:60"`
	SeriShChar    string `json:"seri_sh_char" gorm:"size:100"`
	SeriSh        string `json:"seri_sh" gorm:"size:100"`
	Serial        string `json:"serial" gorm:"size:100"`
	ShNumber      string `json:"sh_number" gorm:"size:100"`
	BirthDate     string `json:"birth_date" gorm:"size:100"`
	PlaceOfIssue  string `json:"place_of_issue" gorm:"size:100"`
	SignatureFile string `json:"signature_file" gorm:"type:text"`
	CustomerId    uint   `json:"customer" gorm:"unique"`
	Gender        string `json:"gender" gorm:"Column:gender;Null"`
}

func (cp *CustomerPrivate) TableName() string {
	return "customer_private_info"
}

func (cp CustomerPrivate) Set(customerPrivate *CustomerPrivate) error {

	db, err := Connect()
	if err != nil {
		return err
	}

	sqlDB := db.DB()
	for {
		if e := sqlDB.Ping(); e == nil {
			break
		}
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(300)

	sqlDB.SetConnMaxLifetime(time.Hour)

	defer db.Close()
	customer := CustomerPrivate{}
	db.Find(&customer, "customer_id=?", customerPrivate.CustomerId)

	tx := db.Begin()

	if customer.ID == 0 {
		if err := tx.Create(&customerPrivate).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}

func (cp CustomerPrivate) GetOrCreate(d CustomerPrivate) (*CustomerPrivate, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}

	private := CustomerPrivate{}
	db.Find(&private, "customer_id", d.CustomerId)

	if private.ID == 0 {
		private = CustomerPrivate{
			FirstName:     d.FirstName,
			LastName:      d.FirstName,
			FatherName:    d.FatherName,
			SeriShChar:    d.SeriShChar,
			SeriSh:        d.SeriSh,
			Serial:        d.Serial,
			ShNumber:      d.ShNumber,
			BirthDate:     d.BirthDate,
			PlaceOfIssue:  d.PlaceOfIssue,
			SignatureFile: d.SignatureFile,
			CustomerId:    d.CustomerId,
			Gender:        d.Gender,
		}

		if err := tx.Create(&private).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		return &private, nil
	}
	return nil, nil

}
