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
