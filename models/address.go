package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CustomerAddress struct {
	gorm.Model
	Address                   string `json:"address" gorm:"Column:address;size:250"`
	PostalCode                string `json:"postal_code" gorm:"Column:postal_code;size:20;Null"`
	CustomerId                uint   `json:"customer_id" gorm:"Column:customer_id"`
	MebbcoType                string `json:"mebbco_type gorm:"size:60;Column:mebbco_type"`
	CountryId                 int32  `json:"country_id" gorm:"Column:country_id;null"`
	CountryName               string `json:"country_name" gorm:"Column:country_name;null;size:100"`
	ProvinceId                int32  `json:"prvince_id" gorm:"Column:province_id; null"`
	ProvinceName              string `json:"province_name" gorm:"Column:province_name;null;size:100"`
	CityId                    int32  `json:"city_id" gorm:"Column:city_id;null"`
	CityName                  string `json:"city_name" gorm:"Column:city_name;size:100"`
	SectionId                 int32  `json:"section_id" gorm:"Column:section_id"`
	SectionName               string `json:"section_name" gorm:"Column:section_name;size:100;null"`
	CityPrefix                string `json: "city_prefix" gorm:"Column:city_prefix;null;size:100"`
	Alley                     string `json: "alley" gorm:"Column:alley;null;size:100"`
	Plaque                    string `json: "plaque" gorm:"Column:plaque;null;size:100"`
	Tel                       string `json:"tel" gorm:"Column:tel;null;size:100"`
	CountryPrefix             string `json:"country_prefix" gorm:"Column:country_prefix;null;size:100"`
	Mobile                    string `json:"mobile" gorm:"Column:mobile;null;size:100"`
	EmergencyTel              string `json:"emergency_tel" gorm:"Column:emergency_tel;null;size:100"`
	EmergencyTelCityPrefix    string `json:"emergency_tel_ciity_prefix" gorm:"null;size:100"`
	EmergencyTelCountryPrefix string `json:"emergency_tel_country_prefix" gorm:"null;size:100"`
	FaxPrefix                 string `json:"fax_prefix" gorm:"null;size:100"`
	Fax                       string `json:"fax" gorm:"null;size:100"`
	Website                   string `json:"website" gorm:"null;size:100"`
	Email                     string `json:"email" gorm:"null;size:100"`
}

func (ca *CustomerAddress) TableName() string {
	return "customer_address"
}

func (ca CustomerAddress) Set(data AddressParam) error {
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
	customerAddress := CustomerAddress{}
	db.Find(&customerAddress, "address=? or postal_code=?", data.Address, data.PostalCode)

	tx := db.Begin()

	if customerAddress.ID == 0 {

		if err := tx.Create(
			&CustomerAddress{
				Address:    data.Address,
				PostalCode: data.PostalCode,
				CustomerId: data.CustomerId,
				MebbcoType: data.MebbcoType,
			},
		).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	return err
}

func (ca CustomerAddress) GetOrCreate(d CustomerAddress) (*CustomerAddress, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	customerAddress := CustomerAddress{}
	db.Find(&customerAddress, "customer_id=?", d.CustomerId)

	if customerAddress.ID == 0 {
		customerAddress = CustomerAddress{
			Address:                   d.Address,
			PostalCode:                d.PostalCode,
			CustomerId:                d.CustomerId,
			MebbcoType:                d.MebbcoType,
			CountryId:                 d.CountryId,
			CountryName:               d.CountryName,
			ProvinceId:                d.ProvinceId,
			ProvinceName:              d.ProvinceName,
			CityId:                    d.CityId,
			CityName:                  d.CityName,
			SectionId:                 d.SectionId,
			SectionName:               d.SectionName,
			CityPrefix:                d.CityPrefix,
			Alley:                     d.Alley,
			Plaque:                    d.Plaque,
			Tel:                       d.Tel,
			CountryPrefix:             d.CountryPrefix,
			Mobile:                    d.Mobile,
			EmergencyTel:              d.EmergencyTel,
			EmergencyTelCityPrefix:    d.EmergencyTelCityPrefix,
			EmergencyTelCountryPrefix: d.EmergencyTelCountryPrefix,
			FaxPrefix:                 d.FaxPrefix,
			Fax:                       d.Fax,
			Website:                   d.Website,
			Email:                     d.Email,
		}

		db.Create(&customerAddress)
	}
	return &customerAddress, nil

}
