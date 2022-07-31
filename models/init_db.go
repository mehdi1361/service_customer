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
		&FinancialInfo{},
		&State{},
		&CustomerState{},
		&FileType{},
		&CustomerFile{},
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
	db.Model(&FinancialInfo{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerState{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerState{}).AddForeignKey("state_id", "base_state(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerFile{}).AddForeignKey("file_type_id", "base_file_type(id)", "CASCADE", "CASCADE")
	db.Model(&CustomerFile{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")

	/*
	var states = []State{
		{StateName: "person_info", Title: "اطلاعات هویتی", IconClass: "fas fa-briefcase", StateId: 1},
		{StateName: "contact_us", Title: "اطلاعات تماس", IconClass: "fas fa-book-reader", StateId: 2},
		{StateName: "account_info", Title: "اطلاعات حساب", IconClass: "fas fa-people-arrows", StateId: 3},
		{StateName: "job_experience", Title: "اطلاعات کاری", IconClass: "fas fa-users", StateId: 4},
		{StateName: "asset", Title: "اطلاعات دارایی", IconClass: "far fa-comment-alt", StateId: 5},
		{StateName: "contract", Title: "بارگزاری", IconClass: "fas fa-images", StateId: 6},
	}
	for _, state := range states {
		db.Create(&state)
	}

	customer := Customer{}
	db.Find(&customer, "id=2")
	db.Create(&FinancialInfo{
		AssetsValue: 10,
		InComingAverage: 100,
		SExchangeTransaction:2000,
		CExchangeTransaction: 2000,
		OutExchangeTransaction: 4000,
		TransactionLevel: "good",
		TradingKnowledgeLevel: "good",
		CompanyPurpose:"test",
		ReferenceRateCompany: "100",
		RateDate:"2012",
		Rate: "12",
		CustomerId: 2,
	})
	db.Create(&Job{Title: "کارمند", JobId: "1"})
	db.Create(&JobInfo{
			EmploymentDate:    "2022-01-01",
			CompanyName:       "خاورمیانه",
			CompanyAddress:    "تهران",
			CompanyPostalCode: "021",
			CompanyEmail:      "mhd.mosavi@gmail.com",
			CompanyWebSite:    "www.test.com",
			CompanyCityPrefix: "021",
			CompanyPhone:      "44522545",
			Position:          "1",
			CompanyFaxPrefix:  "021",
			CompanyFax:        "2121236545",
			JobId:             uint(1),
			JobTitle:          "برنامه نویس",
			JobDescription:    "برنامه نویسی",
			CustomerId:    uint(2),
		})
	var types = []FileType {
		{Name: "signature", FaName: "نمونه امضا", IsActive: true, IsForce: true},
		{Name: "birth_cert_copy", FaName: "کپی شناسنامه", IsActive: true, IsForce: true},
		{Name: "nationa_security_card_copy", FaName: "کپی کارت ملی", IsActive: true, IsForce: true},
		{Name: "etc", FaName: "غیره", IsActive: true, IsForce: false},
	}

	for _, sType := range types {
		db.Create(&sType)
	}
	*/
}
