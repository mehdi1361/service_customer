package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Gender struct {
	gorm.Model
	Name            string `json:"name" gorm:"size:10;unique"`
	CustomerPrivate []CustomerPrivate
}

func (g *Gender) TableName() string {
	return "base_gender"
}

type Country struct {
	gorm.Model
	Name      string     `json:"name" gorm:"size:10;unique"`
	CountryId uint       `json="country_id" gorm:"Column:country_id"`
	Provinces []Province `json="provinces"`
}

func (p *Country) TableName() string {
	return "base_countries"
}

type Province struct {
	gorm.Model
	Name       string `json:"name" gorm:"size:10;unique"`
	ProvinceId uint   `json="coun_id" gorm:"Column:country_id"`
	CountryId  uint   `json="country_id" gorm:"Column:country_id"`
	Cities     []City `json="cities"`
}

func (p *Province) TableName() string {
	return "base_province"
}

type City struct {
	gorm.Model
	Name       string       `json:"name" gorm:"size:10;unique"`
	ProvinceId uint         `json="province_id" gorm:"Column:province_id"`
	CityId     uint         `json="city_id" gorm:"Column:city_id"`
	BankBranch []BankBranch `json:"bank_branch"`
}

func (c *City) TableName() string {
	return "base_cities"
}

type Bank struct {
	gorm.Model
	Title      string       `json:"title" gorm:"size:10"`
	BankId     uint         `json="bank_id" gorm:"Column:bank_id"`
	BankBranch []BankBranch `json:"bank_branch"`
}

func (b *Bank) TableName() string {
	return "base_bank"
}

type BankBranch struct {
	gorm.Model
	Name         string        `json:"name" gorm:"size:10;unique"`
	BankId       uint          `gorm:"Column:bank_id"`
	BankBranchId uint          `gorm:"Column:bank_branch_id"`
	CityId       uint          `gorm:"Column:city_id"`
	BankAccount  []BankAccount `json:"bank_accounts"`
}

func (b *BankBranch) TableName() string {
	return "base_bank_branch"
}

type Domain struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:10;unique"`
	IdRayan int32  `json:"id_rayan"`
}

func (d *Domain) TableName() string {
	return "base_mebbco_domain"
}

type CustomerGroup struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:10;unique"`
	IdRayan int32  `json:"id_rayan"`
}

func (cg *CustomerGroup) TableName() string {
	return "base_mebbco_customer_group"
}

type MebbcoBranch struct {
	gorm.Model
	Title         string         `json:"title" gorm:"size:10;unique"`
	TypeMebbco    string         `json:"type_mebbco" gorm:"size:50;column:type_mebbco"`
	IdRayan       int32          `json:"id_rayan"`
	Portfos       []Portfo       `gorm:"foreignKey:TypeMebbcoId"`
	ComexVisitors []ComexVisitor `gorm:"foreignKey:TypeMebbcoId"`
	TradingCodes []TradingCode `gorm:"foreignKey:TypeMebbcoId"`
}

func (mb *MebbcoBranch) TableName() string {
	return "base_mebbco_branch"
}

type Job struct {
	gorm.Model
	Title string `json:"title" gorm:"size:10;unique"`
	JobId string `json:"job_id" gorm:"size:10"`
	JobInfos []JobInfo `gorm:"foreignKey:JobId"`
}

func (mb *Job) TableName() string {
	return "base_job"
}

type TradingCode struct {
	gorm.Model
	Type string `json:"Type" gorm:"size:50;Column:type"`
	FirstPart string `json:"first_part" gorm:"size:70;Column:first_part"`
	SecondPart string `json:"second_part" gorm:"size:70;Column:second_part"`
	Code string `json:"code" gorm:"size:70;Column:code"`
	IsDefault bool `json:"is_default" gorm:"Column:is_default"`
}

func (tc *TradingCode) TableName() string {
	return "base_job"
}
