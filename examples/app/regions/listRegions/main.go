package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Return the support regions
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Regions(&app.RegionsRequest{})
	fmt.Println(rsp, err)
}
