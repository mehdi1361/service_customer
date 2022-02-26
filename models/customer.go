package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/qor/validations"
	"os"
)

type Customer struct {
	gorm.Model
	SejamReferenceCode string          `json:"sejam_reference_code" gorm:"size:100"`
	UserName           string          `json:"user_name" gorm:"size:100"`
	Password           string          `json:"user_name" gorm:"size:100"`
	IsActive           bool            `json:"is_active"`
	IsRayanService     bool            `json:"is_rayan_service"`
	PhonePerson        []PhonePerson   `json:"phone_persons"`
	PrivateInfo        CustomerPrivate `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
	LegalInfo          CustomerLegal   `gorm:"foreignKey:CustomerId;association_foreignkey:ID"`
}

func (c *Customer) TableName() string {
	return "customer_customer"
}

type CustomerPrivate struct {
	gorm.Model

	FirstName     string `json:"first_name" gorm:"size:60"`
	LastName      string `json:"last_name" gorm:"size:60"`
	FatherName    string `json:"father_name" gorm:"size:60"`
	SeriShChar    string `json:"seri_sh_char" gorm:"size:100"`
	SeriSh        string `json:"seri_sh" gorm:"size:100"`
	Serial        string `json:"serial" gorm:"size:100"`
	ShNumber      string `json:"sh_number" gorm:"size:100"`
	BirthDate     string `json:"birth_date" gorm:"size:100"`
	PlaceOfIssue  string `json:"place_of_issue" gorm:"size:100"`
	SignatureFile string `json:"signature_file" gorm:"type:text"`
	CustomerId    uint   `json:"customer" `
}

func (cp *CustomerPrivate) TableName() string {
	return "customer_private_info"
}

type CustomerLegal struct {
	gorm.Model

	CompanyName            string `json:"company_name" gorm:"size:60"`
	RegisterNumber         string `json:"register_number" gorm:"size:60"`
	RegisterPlace          string `json:"register_place" gorm:"size:60"`
	RegisterDate           string `json:"register_date" gorm:"size:100"`
	EconomicCode           string `json:"economic_code" gorm:"size:100"`
	EvidenceReleaseCompany string `json:"evidence_release_company" gorm:"size:100"`
	EvidenceReleaseDate    string `json:"evidence_release_date" gorm:"size:100"`
	EvidenceExpirationDate string `json:"birth_date" gorm:"size:100"`
	CustomerId             uint   `json:"customer" `
}

func (cl *CustomerLegal) TableName() string {
	return "customer_legal_info"
}

type BankAccount struct {
	gorm.Model
	AccountNumber string `json:"account_number gorm:"size:60"`
	Sheba         string `json:"sheba" gorm:"size:60"`
	IsDefault     bool   `json:"is_default" gorm:"Column:is_default"`
	BranchId      uint   `json:"branch_id" gorm:"Column:branch_id"`
}

func (ba *BankAccount) TableName() string {
	return "customer_bank_account"
}

type PhonePerson struct {
	gorm.Model
	AccountNumber string `json:"account_number gorm:"size:60"`
	Sheba         string `json:"sheba" gorm:"size:60"`
	IsActive      bool   `json:"is_active" gorm:"Column:is_active"`
	CustomerId    uint   `json:"customer_id" gorm:"Column:customer_id"`
}

func (pp *PhonePerson) TableName() string {
	return "customer_phone_person"
}

type FinancialInfo struct {
	AssetValue             int32 `json:"asset_value" gorm:"Column:asset_value"`
	InCommingAverage       int32 `json:"incoming_average" gorm:"Column:asset_value"`
	SExchangeTransaction   int32 `json:"s_exchange_tranasction" gorm:"Column:s_exchange_tranasction"`
	CExchangeTransaction   int32 `json:"c_exchange_tranasction" gorm:"Column:c_exchange_tranasction"`
	OutExchangeTransaction int32 `json:"out_exchange_tranasction" gorm:"Column:out_exchange_tranasction"`
	TransactionLevel       int32 `json:"tranasction_level" gorm:"Column:tranasction_level"`
	TradingKnowledeLevel       int32 `json:"trading_knowledge_level" gorm:"Column:trading_knowledge_level"`
}

func (fi *FinancialInfo) TableName() string {
	return "customer_financial_info"
}

func Connect() (db *gorm.DB, err error) {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	server := os.Getenv("DB_SERVER")
	database := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	conn, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", server, port, dbUser, database, password),
	)
	validations.RegisterCallbacks(conn)
	return conn, err
}

func init() {
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
	)
	db.Model(&Province{}).AddForeignKey("country_id", "base_countries(id)", "CASCADE", "CASCADE")
	db.Model(&City{}).AddForeignKey("province_id", "base_province(id)", "CASCADE", "CASCADE")
	db.Model(&BankBranch{}).AddForeignKey("bank_id", "base_bank(id)", "CASCADE", "CASCADE")
	db.Model(&BankBranch{}).AddForeignKey("city_id", "base_cities(id)", "CASCADE", "CASCADE")
	db.Model(&BankAccount{}).AddForeignKey("branch_id", "base_bank_branch(id)", "CASCADE", "CASCADE")
	db.Model(&PhonePerson{}).AddForeignKey("customer_id", "customer_customer(id)", "CASCADE", "CASCADE")
	//	db.Model(&CustomerPrivate{}).AddForeignKey("id", "customer_customer(id)", "CASCADE", "CASCADE")
}
