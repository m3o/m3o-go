package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Update the app. The latest source code will be downloaded, built and deployed.
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Update(&app.UpdateRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
