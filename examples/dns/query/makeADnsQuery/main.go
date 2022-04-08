package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/dns"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Dns.Query(&dns.QueryRequest{
		Name: "google.com",
	})
	fmt.Println(rsp, err)
}
