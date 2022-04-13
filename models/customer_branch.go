package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CustomerBranch struct {
	gorm.Model
	CustomerId uint `json:"customer" gorm:"Column:customer_id" sql:"unique_index:idx_branch_customer"`
	BranchId   uint `json:"customer" gorm:"Column:branch_id" sql:"unique_index:idx_branch_customer"`
	MebbcoType  string `json:"mebbco_type gorm:"size:60;Column:mebbco_type"`
}

func(cb *CustomerBranch) TableName() string {
	return "customer_branch"
}


func (cb CustomerBranch) Set(customerId uint, typeName string, branch *MebbcoBranch) error {

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
	err, findBranch := MebbcoBranch{}.Set(branch)
	if err != nil {
	 return err
	}

	defer db.Close()
	customerb := CustomerBranch{}
	db.Find(&customerb, "branch_id=? and customer_id=?", findBranch.ID, customerId)

	tx := db.Begin()

	if customerb.ID == 0 {
		if err := tx.Create(&CustomerBranch{CustomerId: customerId, BranchId: findBranch.ID, MebbcoType: typeName}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}
