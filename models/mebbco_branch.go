package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MebbcoBranch struct {
	gorm.Model
	Title         string         `json:"title" gorm:"size:100"`
	TypeMebbco    string         `json:"type_mebbco" gorm:"size:50;column:type_mebbco"`
	IdRayan       int64            `json:"id_rayan" gorm:"unique;Column:id_rayan"`
	Portfos       []Portfo       `gorm:"foreignKey:TypeMebbcoId"`
	ComexVisitors []ComexVisitor `gorm:"foreignKey:TypeMebbcoId"`
	TradingCodes  []TradingCode  `gorm:"foreignKey:TypeMebbcoId"`
}

func (mb *MebbcoBranch) TableName() string {
	return "base_mebbco_branch"
}

func (mb MebbcoBranch) Set(mebbcoBranch *MebbcoBranch) (error, *MebbcoBranch) {

	db, err := Connect()
	if err != nil {
		return err, nil
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
	branch := MebbcoBranch{}
	db.Find(&branch, "id_rayan=?", mebbcoBranch.IdRayan)

	if branch.ID == 0 {
		db.Create(mebbcoBranch)
		return nil, mebbcoBranch

	} else {
		return nil, &branch
	}

}
