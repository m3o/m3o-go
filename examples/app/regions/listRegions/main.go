package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Return the support regions
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Regions(&app.RegionsRequest{})
	fmt.Println(rsp, err)
}
