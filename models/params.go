package models

type SetCityParams struct {
	ProvinceId   int
	ProvinceName string
	CityId       int
	CityName     string
}

type ProvinceParam struct {
	ProvinceId   int
	ProvinceName string
}

type AddressParam struct {
	Address    string
	PostalCode string
	CustomerId uint
	MebbcoType string
}

type BourseAccountParam struct {
	Id         int64
	Name       string
	Number     string
	IsDefault  int64
	CustomerId uint
}

type CustomerGroupParams struct {
	Title      string
	RayanId    int64
	CustomerId uint
}

type DomainParams struct {
	Title      string
	RayanId    int64
	CustomerId uint
}

type BankParam struct {
	BankId   int64
	BankName string
}

type BankBranchParam struct {
	BankId         int64
	BankName       string
	BankBranchCode string
	Name           string
}

type BankAccountParam struct {
	AccountNumber      string
	RayanBankAccountId int64
	BaTypeName         string
	Shaba              string
	IsDefault          int64
	IsActive           bool
	IsOnline           int64
	BankId             int64
	BankName           string
	BankBranchCode     string
	BranchName         string
	CustomerId         uint
}
