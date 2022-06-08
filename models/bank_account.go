package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BankAccount struct {
	gorm.Model
	AccountNumber      string `json:"account_number gorm:"size:60"`
	RayanBankAccountId int64  `json:"rayan_bank_account_id" gorm: "Column:rayan_bank_account_id"`
	BaTypeName         string `json:"sheba" gorm:"size:60"`
	Shaba              string `json:"shaba" gorm:"size:60"`
	IsDefault          int64  `json:"is_default" gorm:"Column:is_default"`
	IsActive           bool   `json:"is_active" gorm:"Column:is_active"`
	IsOnline           int64  `json:"is_online" gorm:"Column:is_online"`
	BranchId           uint   `json:"branch_id" gorm:"Column:branch_id"`
	CustomerId         uint   `json:"customer_id" gorm:"Column:customer_id"`
}

func (ba *BankAccount) TableName() string {
	return "customer_bank_account"
}

func (ba BankAccount) GetOrCreate(data BankAccountParam) (*BankAccount, error) {
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

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(300)

	sqlDB.SetConnMaxLifetime(time.Hour)

	defer db.Close()
	bankBranch, err := BankBranch{}.GetOrCreate(BankBranchParam{
		BankId:         data.BankId,
		BankName:       data.BankName,
		BankBranchCode: data.BankBranchCode,
		Name:           data.BranchName,
	})

	if err != nil {
		return nil, err
	}
	bankAccount := &BankAccount{}
	db.Find(&bankAccount, "shaba=? and customer_id=?", data.Shaba, data.CustomerId)

	tx := db.Begin()

	if bankAccount.ID == 0 {
		if err := tx.Create(
			&BankAccount{
				AccountNumber:      data.AccountNumber,
				RayanBankAccountId: data.RayanBankAccountId,
				BaTypeName:         data.BaTypeName,
				Shaba:              data.Shaba,
				IsDefault:          data.IsDefault,
				IsActive:           data.IsActive,
				IsOnline:           data.IsOnline,
				BranchId:           bankBranch.ID,
				CustomerId:         data.CustomerId,
			},
		).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		db.Find(&bankAccount, "shaba=? and customer_id=?", data.Shaba, data.CustomerId)
	}

	err = tx.Commit().Error
	return bankAccount, err
}
