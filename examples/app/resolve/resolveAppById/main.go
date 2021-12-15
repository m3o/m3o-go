package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Resolve an app by id to its raw backend endpoint
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Resolve(&app.ResolveRequest{
		Id: "helloworld",
	})
	fmt.Println(rsp, err)
}
