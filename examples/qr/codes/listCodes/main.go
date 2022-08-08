package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/qr"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Qr.Codes(&qr.CodesRequest{})
	fmt.Println(rsp, err)
}
