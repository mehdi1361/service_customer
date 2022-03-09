package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	service "service_customer/service/proto"
	_ "service_customer/scheduler"
	gServer "service_customer/server"
)

func main() {
	grpcServer := grpc.NewServer()
	var server gServer.Server
	service.RegisterCustomerServiceServer(grpcServer, server)
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}
