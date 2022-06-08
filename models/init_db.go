package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
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
		&CustomerBranch{},
		&VerificationCode{},
		&Broker{},
		&CustomerAddress{},
		&BourseAccounts{},
	)
	db.Model(&Province{}).AddForeignKey("country_id", "base_countries(id)", "CASCADE", "CASCADE")
	db.Model(&City{}).AddForeignKey("province_id", "base_province(id)", "CASCADE", "CASCADE")
	db.Model(&BankBranch{}).AddForeignKey("bank_id", "base_bank(id)", "CASCADE", "CASCADE")
	//	db.Model(&BankBranch{}).AddForeignKey("city_id", "base_cities(id)", "CASCADE", "CASCADE")
	db.Model(&BankAccount{}).AddForeignKey("branch_id", "base_bank_branch(id)", "CASCADE", "CASCADE")
	db.Model(&PhonePerson{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&Portfo{}).AddForeignKey("type_mebbco_id", "base_mebbco_branch(id)", "CASCADE", "CASCADE")
	db.Model(&ComexVisitor{}).AddForeignKey("type_mebbco_id", "base_mebbco_branch(id)", "CASCADE", "CASCADE")
	db.Model(&JobInfo{}).AddForeignKey("job_id", "base_job(id)", "CASCADE", "CASCADE")
	db.Model(&JobInfo{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&Fund{}).AddForeignKey("customer_service_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerPrivate{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerLegal{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerBranch{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerBranch{}).AddForeignKey("branch_id", "base_mebbco_branch(id)", "CASCADE", "CASCADE")
	db.Model(&VerificationCode{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&Broker{}).AddForeignKey("customer_service_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerAddress{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&BourseAccounts{}).AddForeignKey("customer_id", "customer_broker_info(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerGroup{}).AddForeignKey("customer_id", "customer_broker_info(id)", "CASCADE", "CASCADE")
	db.Model(&BankAccount{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
}
