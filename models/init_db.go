package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	fmt.Println("test dbbbbbbbbbbbbbbbbbbbbbb")
	conn, err := Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()

	db := conn

	_ = db.AutoMigrate(
		&Customer{},
		&CustomerPrivate{},
		&CustomerLegal{},
		&Gender{},
		&Country{},
		&Province{},
		&City{},
		&Bank{},
		&BankBranch{},
		&BankAccount{},
		&Domain{},
		&CustomerGroup{},
		&MebbcoBranch{},
		&PhonePerson{},
		&Job{},
		&Portfo{},
		&ComexVisitor{},
		&JobInfo{},
		&Fund{},
	)
	db.Model(&Province{}).AddForeignKey("country_id", "base_countries(id)", "CASCADE", "CASCADE")
	db.Model(&City{}).AddForeignKey("province_id", "base_province(id)", "CASCADE", "CASCADE")
	db.Model(&BankBranch{}).AddForeignKey("bank_id", "base_bank(id)", "CASCADE", "CASCADE")
	db.Model(&BankBranch{}).AddForeignKey("city_id", "base_cities(id)", "CASCADE", "CASCADE")
	db.Model(&BankAccount{}).AddForeignKey("branch_id", "base_bank_branch(id)", "CASCADE", "CASCADE")
	db.Model(&PhonePerson{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&Portfo{}).AddForeignKey("type_mebbco_id", "base_mebbco_branch(id)", "CASCADE", "CASCADE")
	db.Model(&ComexVisitor{}).AddForeignKey("type_mebbco_id", "base_mebbco_branch(id)", "CASCADE", "CASCADE")
	db.Model(&JobInfo{}).AddForeignKey("job_id", "base_job(id)", "CASCADE", "CASCADE")
	db.Model(&JobInfo{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&Fund{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerPrivate{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerLegal{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
}
