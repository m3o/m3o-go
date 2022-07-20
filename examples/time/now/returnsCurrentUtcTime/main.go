package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/time"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Time.Now(&time.NowRequest{})
	fmt.Println(rsp, err)
}
