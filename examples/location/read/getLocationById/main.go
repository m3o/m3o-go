package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/location"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Location.Read(&location.ReadRequest{
		Id: "1",
	})
	fmt.Println(rsp, err)
}
