package scheduler

import (
	"fmt"
	"github.com/k0kubun/pp"
	"golang.org/x/net/context"
	"log"
	models "service_customer/models"
	service "service_customer/service/rayan/proto"
	//"strconv"
	//"strings"
)

func readCustomerBrokerInfo() {

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
			log.Printf("error: %s", err)
			return
		} else {
			pp.Print(customer.Result[0])
		}

		go func() {
			err = models.PhonePerson{}.SetPhone(&models.PhonePerson{
				PhoneNumber: customer.Result[0].CellPhone,
				IsMobile:    true,
				MebbcoType:  "broker",
				CustomerId:  v.CustomerServiceId,
			})

			err = models.PhonePerson{}.SetPhone(&models.PhonePerson{
				PhoneNumber: customer.Result[0].Phone,
				IsMobile:    true,
				MebbcoType:  "broker",
				CustomerId:  v.CustomerServiceId,
			})
		}()

		go func() {
			if customer.Result[0].IsLegal == 0 {
				err = models.CustomerPrivate{}.Set(&models.CustomerPrivate{
					FirstName:  customer.Result[0].FirstName,
					LastName:   customer.Result[0].LastName,
					FatherName: customer.Result[0].Parent,
					ShNumber:   customer.Result[0].BirthCertificationNumber,
					SeriSh:     customer.Result[0].BirthCertId,
					CustomerId: v.CustomerServiceId,
				})
			} else {

				err = models.CustomerLegal{}.Set(&models.CustomerLegal{
					CompanyName:    customer.Result[0].CompanyName,
					RegisterNumber: customer.Result[0].RegistrationNumber,
					CustomerId:     v.CustomerServiceId,
				})
			}
		}()

		err = models.CustomerBranch{}.Set(
			v.CustomerServiceId,
			"broker",
			&models.MebbcoBranch{
				Title:   customer.Result[0].BranchName,
				IdRayan: customer.Result[0].BranchId,
			})

		go func() {
			for _, value := range customer.Result[0].BourseAccounts {
				models.BourseAccounts{}.GetOrCreate(
					models.BourseAccountParam{
						Id:     value.BourseAccountId,
						Name:   value.BourseAccountName,
						Number: value.AccountNumber,
						IsDefault: value.IsDefault,
						CustomerId: v.ID,
					},
				)
			}
		}()
	}

}
