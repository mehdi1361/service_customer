package scheduler

import (
	"github.com/jasonlvhit/gocron"
)

func init() {
	//readCustomerFundData()
	go func() {
		//		gocron.Every(60).Second().Do(readCustomerFundData)
		gocron.Every(60).Second().Do(readCustomerFundInfo)
		<-gocron.Start()
	}()
}
