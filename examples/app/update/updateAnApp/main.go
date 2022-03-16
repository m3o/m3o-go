package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Update the app. The latest source code will be downloaded, built and deployed.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Update(&app.UpdateRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
