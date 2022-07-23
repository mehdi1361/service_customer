package models

import (
	"github.com/jinzhu/gorm"
)

type JobInfo struct {
	gorm.Model
	EmploymentDate    string `json:"employment_date" gorm:"Column:employment_date;size:60"`
	CompanyName       string `json:"company_name" gorm:"Column:company_name;size:60"`
	CompanyAddress    string `json:"company_address" gorm:"Column:company_address;size:60"`
	CompanyPostalCode string `json:"company_postal_code" gorm:"Column:company_postal_code;size:60"`
	CompanyEmail      string `json:"company_email" gorm:"Column:company_email;size:60"`
	CompanyWebSite    string `json:"company_website" gorm:"Column:company_website;size:60"`
	CompanyCityPrefix string `json:"company_city_prefix" gorm:"Column:company_city_prefix;size:60"`
	CompanyPhone      string `json:"company_phone" gorm:"Column:company_phone;size:60"`
	JobDescription    string `json:"job_description" gorm:"Column:job_description;size:60"`
	Position          string `json:"position" gorm:"Column:position;size:60"`
	CompanyFaxPrefix  string `json:"company_fax_prefix" gorm:"Column:company_fax_prefix;size:60"`
	CompanyFax        string `json:"company_fax" gorm:"Column:company_fax;size:60"`
	CustomerId        uint   `gorm:"Column:customer_id"`
	JobId             uint   `gorm:"Column:job_id"`
	JobTitle          string `json:"job_title" gorm:"Column:job_title;size:60"`
}

func (j JobInfo) TableName() string {
	return "customer_job_info"
}

func (j JobInfo) GetOrCreate(d JobInfo) (*JobInfo, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	job := JobInfo{}
	db.Find(&job, "customer_id=?", d.CustomerId)

	if job.ID == 0 {
		job = JobInfo{
			EmploymentDate:    d.EmploymentDate,
			CompanyName:       d.CompanyName,
			CompanyAddress:    d.CompanyAddress,
			CompanyPostalCode: d.CompanyPostalCode,
			CompanyEmail:      d.CompanyEmail,
			CompanyWebSite:    d.CompanyWebSite,
			CompanyCityPrefix: d.CompanyCityPrefix,
			CompanyPhone:      d.CompanyPhone,
			Position:          d.Position,
			CompanyFaxPrefix:  d.CompanyFaxPrefix,
			CompanyFax:        d.CompanyFax,
			JobId:             d.JobId,
			JobTitle:          d.JobTitle,
			JobDescription:    d.JobDescription,
		}

		db.Create(&job)

	}
		return &job, nil

}
