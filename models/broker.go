package models

import (
	"github.com/jinzhu/gorm"
)

type Broker struct {
	gorm.Model
	TelegramUsername           string `json:"telegram_username" gorm:"Column:telegram_username;size:100"`
	TelegramStatusId           string `json:"telegram_status_id" gorm:"Column:telegram_status_id;size:100"`
	RegisterationNumber        string `json:"registeration_number" gorm:"Column:registeration_number;size:100"`
	BourseAccountName          string `json:"bourse_account_name" gorm:"Column:bourse_account_name;size:100"`
	AccountNumber              string `json:"account_number" gorm:"Column:account_number;size:100"`
	OnlineUsername             string `json:"online_username" gorm:"Column:online_username;size:100"`
	HasOnlineAccount           int64  `json:"has_online_account" gorm:"Column:has_online_account;size:100"`
	ModificationDate           string `json:"modification_date" gorm:"Column:modification_date;size:100"`
	IsMmtpUser                 int64  `json:"is_mmtp_user" gorm:"Column:is_mmtp_user"`
	MmtpUserStatusId           int64  `json:"mmtp_user_status_id" gorm:"Column:mmtp_user_status_id"`
	IsSiteUser                 int64  `json:"is_site_user" gorm:"Column:is_site_user"`
	EorderStatusId             int64  `json:"e_order_status_id" gorm:"Column:e_order_status_id"`
	HasSignSample              bool   `json:"has_sign_sample" gorm:"Column:has_sign_sample"`
	HasCustomerPhoto           bool   `json:"has_customer_photo" gorm:"Column:has_customer_photo"`
	HasBirthCertificate        bool   `json:"has_birth_certificate" gorm:"Column:has_birth_certificate"`
	HasCertificateComments     bool   `json:"has_certificate_comments" gorm:"Column:has_certificate_comments"`
	HasZipFile                 bool   `json:"has_zip_file" gorm:"Column:has_zip_file"`
	HasOfficialGazette         bool   `json:"has_official_gazette" gorm:"Column:has_official_gazette"`
	HasOfficialAds             bool   `json:"has_official_ads" gorm:"Column:has_official_ads"`
	ComexVisitorId             int64  `json:"comex_visitor_id" gorm:"Column:comex_visitor_id"`
	MmtpUserId                 int64  `json:"mmtp_user_id" gorm:"Column:mmtp_user_id"`
	ComexEconomyAccount        string `json:"comex_economy_account" gorm:"Column:comex_economy_account;size:100"`
	IsPortfo                   int64  `json:"is_portfo" gorm:"Column:is_portfo"`
	TraderCredit               int64  `json:"trader_credit" gorm:"Column:trader_credit"`
	ComexCredit                int64  `json:"comex_credit" gorm:"Column:comex_credit"`
	SfCredit                   int64  `json:"sf_credit" gorm:"Column:sf_credit"`
	Credit                     int64  `json:"credit" gorm:"Column:credit"`
	IsStockCreditPurchase      int64  `json:"is_stock_credit_purchase" gorm:"Column:is_stock_credit_purchase"`
	IsCollateralStocksCustomer int64  `json:"is_collateral_stocks_customer" gorm:"Column:is_collateral_stocks_customer"`
	CustomerIdentity           string `json:"customer_identity" gorm:"Column:customer_identity;size:100"`
	CustomerId                 int64  `json:"customer" gorm:"Column:customer_id;size:100"`
	CustomerServiceId          uint   `json:"customer" gorm:"Column:customer_service_id;size:100"`
}

func (b *Broker) TableName() string {
	return "customer_broker_info"
}
