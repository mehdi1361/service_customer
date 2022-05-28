package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	fund_customer_data, _ := strconv.ParseUint(os.Getenv("FUND_CUSTOMER_DATA_RUN_EVERY_SECOND"), 10, 64)
	fund_customer_info, _ := strconv.ParseUint(os.Getenv("FUND_CUSTOMER_INFO_RUN_EVERY_SECOND"), 10, 64)
	broker_customer_data, _ := strconv.ParseUint(os.Getenv("BROKER_CUSTOMER_DATA_RUN_EVERY_SECOND"), 10, 64)

	go func() {
		gocron.Every(fund_customer_data).Second().Do(readCustomerFundData)
		gocron.Every(fund_customer_info).Second().Do(readCustomerFundInfo)
		gocron.Every(broker_customer_data).Second().Do(readBrokerCustomerData)
		<-gocron.Start()
	}()
	readCustomerBrokerInfo()
}
