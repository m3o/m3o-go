package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/lists"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	stream, err := client.Lists.Events(&lists.EventsRequest{
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
