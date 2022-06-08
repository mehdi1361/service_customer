package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bank struct {
	gorm.Model
	Title      string       `json:"title" gorm:"size:100"`
	BankId     int64        `json="bank_id" gorm:"Column:bank_id;unique"`
	BankBranch []BankBranch `json:"bank_branch"`
}

func (b *Bank) TableName() string {
	return "base_bank"
}

func (ba Bank) GetOrCreate(data BankParam) (*Bank, error) {
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
	bank := Bank{}
	db.Find(&bank, "bank_id=?", data.BankId)

	tx := db.Begin()

	if bank.ID == 0 {
		err := tx.Create(
			&Bank{
				BankId: data.BankId,
				Title:  data.BankName,
			},
		).Error

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		db.Find(&bank, "bank_id=?", data.BankId)
	}

	err = tx.Commit().Error
	return &bank, err
}
