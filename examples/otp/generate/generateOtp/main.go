package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/otp"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Otp.Generate(&otp.GenerateRequest{
		Id: "asim@example.com",
	})
	fmt.Println(rsp, err)
}
