package scheduler

import (
	"fmt"
	"github.com/k0kubun/pp"
	"golang.org/x/net/context"
	"log"
	models "service_customer/models"
	service "service_customer/service/rayan/proto"
	"strconv"
	//"strings"
)

func readCustomerFundInfo() {
	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()

	customerAll := []models.Customer{}
	db.Find(&customerAll)

	conn, err := rayanConnection()

	if err != nil {
		log.Fatal("error in connection")
		return
	}

	defer conn.Close()

	for _, v := range customerAll {
		customerFund := []models.Fund{}
		db.Find(&customerFund, "customer_service_id=?", v.ID)

		for _, t := range customerFund {
			c := service.NewRecommendationsClient(conn)
			customer, err := c.FundCustomerInfoService(context.Background(), &service.MainRequest{RayanCustomerId: t.CustomerId, Name: t.FundName})

			if err != nil {
				log.Printf("error: %s", err)
				return
			}
			fmt.Println(customer.Result.IsLegal)

			go func() {
				err = models.PhonePerson{}.SetPhone(&models.PhonePerson{
					PhoneNumber: customer.Result.MobileNumber,
					IsActive:    true,
					IsMobile:    true,
					MebbcoType:  t.FundName,
					CustomerId:  v.ID,
				})
			}()

			if err != nil {
				log.Printf("error: %s customer id %s, %v", err, v.ID, customer)
			}

			go func() {
				err = models.PhonePerson{}.SetPhone(&models.PhonePerson{
					PhoneNumber: customer.Result.Phone,
					IsActive:    false,
					IsMobile:    false,
					MebbcoType:  t.FundName,
					CustomerId:  v.ID,
				})
			}()

			if err != nil {
				log.Printf("error: %s customer id %s, %v", err, v.ID, customer)
				return
			}
			go func() {
				if customer.Result.IsLegal == "0" {
					err = models.CustomerPrivate{}.Set(&models.CustomerPrivate{
						FirstName:  customer.Result.FirstName,
						LastName:   customer.Result.LastName,
						FatherName: customer.Result.FatherName,
						ShNumber:   customer.Result.BirthCertNumber,
						SeriSh:     customer.Result.BirthCertificationId,
						CustomerId: v.ID,
					})
				} else {

					err = models.CustomerLegal{}.Set(&models.CustomerLegal{
						CompanyName:    customer.Result.CompanyName,
						RegisterNumber: customer.Result.RegistrationId,
						CustomerId:     v.ID,
					})
				}
			}()

			idRayan, err := strconv.Atoi(customer.Result.BranchId)

			err = models.CustomerBranch{}.Set(
				v.ID,
				t.FundName,
				&models.MebbcoBranch{
					Title: customer.Result.BranchName,
					IdRayan: idRayan,
				})

			if err == nil {
				pp.Print(customer.Result)
			}

		}

	}
}
