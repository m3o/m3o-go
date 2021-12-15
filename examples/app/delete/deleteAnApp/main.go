package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Delete an app
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Delete(&app.DeleteRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
