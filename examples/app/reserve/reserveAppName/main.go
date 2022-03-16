package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Reserve apps beyond the free quota. Call Run after.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Reserve(&app.ReserveRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
