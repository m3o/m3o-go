package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// List all the apps
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.List(&app.ListRequest{})
	fmt.Println(rsp, err)
}
