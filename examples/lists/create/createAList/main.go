package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/lists"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Lists.Create(&lists.CreateRequest{})
	fmt.Println(rsp, err)
}
