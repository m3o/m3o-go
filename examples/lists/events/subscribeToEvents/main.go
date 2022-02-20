package main

import (
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Subscribe to lists events
func main() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	stream, err := listsService.Events(&lists.EventsRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
