package scheduler

import (
	"golang.org/x/net/context"
	"log"
	models "service_customer/models"
	service "service_customer/service/rayan/proto"
	"strings"
)

func readCustomerFundData() {
	conn, err := rayanConnection()

	if err != nil {
		log.Printf("error in connection")
		return
	}

	defer conn.Close()

	c := service.NewRecommendationsClient(conn)
	response, err := c.RayanAccountService(context.Background(), &service.Request{})

	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	for _, account_data := range response.Result {
		if strings.Contains(account_data.Name, "fund") {
			customers, err := c.FundCustomerListService(context.Background(), &service.MainRequest{Name: account_data.Name})
			if err != nil {
				log.Printf("error: %s", err)
			}
			models.Customer{}.SetBulkDataFund(customers.Result, account_data.Name)
		}
	}
}
