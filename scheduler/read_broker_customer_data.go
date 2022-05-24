package scheduler

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	//models "service_customer/models"
	service "service_customer/service/rayan/proto"
)

func readBrokerCustomerData() {
	conn, err := rayanConnection()

	if err != nil {
		log.Printf("error in connection")
		return
	}

	defer conn.Close()

	c := service.NewRecommendationsClient(conn)
	customers, err := c.BrokerCustomerListService(context.Background(), &service.MainRequest{Name: "broker"})
	if err != nil {
		log.Printf("error: %s", err)
	} else {

		fmt.Println(customers.Result)
	}
	//models.Customer{}.SetBulkDataFund(customers.Result, "broker")
}
