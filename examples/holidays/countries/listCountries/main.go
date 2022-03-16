package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/holidays"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Holidays.Countries(&holidays.CountriesRequest{})
	fmt.Println(rsp, err)
}
