package scheduler

import (
	"github.com/jasonlvhit/gocron"
)

func init() {
	/*	envErr := godotenv.Load()
		if envErr != nil {
			panic(envErr)
		}
	*/
	//readCustomerFundData()
	go func() {
		gocron.Every(60).Second().Do(readCustomerFundData)
		<-gocron.Start()
	}()
}
