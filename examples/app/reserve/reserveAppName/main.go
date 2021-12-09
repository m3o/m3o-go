package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Reserve your app name
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Reserve(&app.ReserveRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
