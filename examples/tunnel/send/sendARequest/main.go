package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/tunnel"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Tunnel.Send(&tunnel.SendRequest{
		Url: "https://google.com",
	})
	fmt.Println(rsp, err)
}
