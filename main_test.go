package main
import (
	empty "github.com/golang/protobuf/ptypes/empty"
  "log"
	"testing"
	"golang.org/x/net/context"
)
func TestCountry(t *testing.T) {
  ctx := context.Background()
	request := empty.Empty{}
  server := Server{}
  response, err := server.All(ctx, &request)
  if err != nil {
    t.Error(err)
  }
  log.Println(response)
}
