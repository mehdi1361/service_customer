package scheduler

import (
	"fmt"
	"google.golang.org/grpc"
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func rayanConnection() (*grpc.ClientConn, error) {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}

	rayanUrl := os.Getenv("RAYAN_URL")
	rayanPort := os.Getenv("RAYAN_PORT")
	maxMsgRecSize, _ := strconv.Atoi(os.Getenv("RAYAN_MAX_RECIEVE_SIZE"))
	maxMsgSize, _ := strconv.Atoi(os.Getenv("RAYAN_MAX_SEND_SIZE"))

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
