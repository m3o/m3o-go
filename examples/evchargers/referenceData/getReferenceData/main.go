package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/evchargers"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Evchargers.ReferenceData(&evchargers.ReferenceDataRequest{})
	fmt.Println(rsp, err)
}
