package service_connections

import (
	"fmt"
	"google.golang.org/grpc"
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func SmsServiceConnection() (*grpc.ClientConn, error) {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}

	rayanUrl := os.Getenv("SMS_SERVICE_URL")
	rayanPort := os.Getenv("SMS_SERVICE_PORT")
	maxMsgRecSize, _ := strconv.Atoi(os.Getenv("SMS_SERVICE_MAX_RECIEVE_SIZE"))
	maxMsgSize, _ := strconv.Atoi(os.Getenv("SMS_SERVICE_MAX_SEND_SIZE"))

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", rayanUrl, rayanPort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgRecSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)

	if err != nil {
		return nil, errors.New("cant connect to rayan service")
	}
	return conn, nil
}
