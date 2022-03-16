package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/qr"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Qr.Generate(&qr.GenerateRequest{
		Size: 300,
		Text: "https://m3o.com/qr",
	})
	fmt.Println(rsp, err)
}
