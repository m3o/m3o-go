package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Get the status of an app
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Status(&app.StatusRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
