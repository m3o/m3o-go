package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/postcode"
)

// Return a random postcode and its related info
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Postcode.Random(&postcode.RandomRequest{})
	fmt.Println(rsp, err)
}
