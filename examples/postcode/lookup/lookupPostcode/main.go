package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/postcode"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Postcode.Lookup(&postcode.LookupRequest{
		Postcode: "SW1A 2AA",
	})
	fmt.Println(rsp, err)
}
