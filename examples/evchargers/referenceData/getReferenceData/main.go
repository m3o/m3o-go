package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/evchargers"
)

// Retrieve reference data as used by this API and in conjunction with the Search endpoint
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Evchargers.ReferenceData(&evchargers.ReferenceDataRequest{})
	fmt.Println(rsp, err)
}
