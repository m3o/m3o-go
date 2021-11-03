package example

import (
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// List messages for a given channel
func ListMessages() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.ListMessages(&stream.ListMessagesRequest{
		Channel: "general",
	})
	fmt.Println(rsp, err)
}
