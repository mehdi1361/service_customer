package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BankBranch struct {
	gorm.Model
	Name        string        `json:"name" gorm:"size:100"`
	Code        string        `json:"code" gorm:"size:100;null"`
	SejamCode   string        `json:"sejam_code" gorm:"size:100;null"`
	DlNumber    string        `json:"dl_number" gorm:"size:100;null"`
	RayanBankId string        `json:"rayan_bank_id" gorm:"Column:rayan_bank_id"`
	BankId      uint          `gorm:"Column:bank_id"`
	CityId      uint          `gorm:"Column:city_id"`
	BankAccount []BankAccount `json:"bank_accounts"`
}

func (b *BankBranch) TableName() string {
	return "base_bank_branch"
}

func (ba BankBranch) GetOrCreate(data BankBranchParam) (*BankBranch, error) {
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
	bank, err := Bank{}.GetOrCreate(BankParam{
		BankId:   data.BankId,
		BankName: data.BankName,
	})

	if err != nil {
		return nil, err
	}
	bankBranch := BankBranch{}
	db.Find(&bankBranch, "name=? and bank_id=?", data.Name, bank.ID)

	tx := db.Begin()

	if bankBranch.ID == 0 {
		if err := tx.Create(
			&BankBranch{
				Name:        data.Name,
				RayanBankId: data.BankBranchCode,
				BankId:      bank.ID,
			},
		).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		db.Find(&bankBranch, "name=? and bank_id=?", data.Name, bank.ID)
	}

	err = tx.Commit().Error
	return &bankBranch, err
}
