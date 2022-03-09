package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	service "service_customer/service/rayan/proto"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.3.4.2:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("error in connection")
	}
	defer conn.Close()

	c := service.NewRecommendationsClient(conn)
	response, err := c.RayanAccountService(context.Background(), &service.Request{})

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(response.Result)
}
