package scheduler

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	models "service_customer/models"
	service "service_customer/service/rayan/proto"
	"strings"
)

func readCustomerFundData() {
	conn, err := rayanConnection()

	if err != nil {
		log.Fatal("error in connection")
	}

	defer conn.Close()

	c := service.NewRecommendationsClient(conn)
	response, err := c.RayanAccountService(context.Background(), &service.Request{})

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	for _, account_data := range response.Result {
		if strings.Contains(account_data.Name, "fund") {
			fmt.Println(account_data.Name)
			customers, err := c.FundCustomerListService(context.Background(), &service.MainRequest{Name: account_data.Name})
			if err != nil {
				log.Fatalf("error: %s", err)
			}
			fmt.Println(len(customers.Result))
			models.Customer{}.SetBulkDataFund(customers.Result)
		}
	}
}
