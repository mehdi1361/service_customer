package scheduler

import (
	"fmt"
	//	"github.com/k0kubun/pp"
	"golang.org/x/net/context"
	"log"
	models "service_customer/models"
	service "service_customer/service/rayan/proto"
	//"strconv"
	//"strings"
)

func readCustomerBrokerInfo() {
	fmt.Println("test")

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()

	conn, err := rayanConnection()

	if err != nil {
		log.Fatal("error in connection")
		return
	}

	defer conn.Close()

	customerBrokerAll := []models.Broker{}
	db.Find(&customerBrokerAll)

	for _, v := range customerBrokerAll {

		c := service.NewRecommendationsClient(conn)
		customer, err := c.BrokerCustomerService(context.Background(), &service.MainRequest{RayanCustomerId: v.CustomerId, Name: "broker"})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(customer.Result[0].CityName, customer.Result[0].ProvinceName)
		}
	}

}
