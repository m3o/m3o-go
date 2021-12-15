package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Run an app
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Run(&app.RunRequest{
		Branch: "master",
		Name:   "helloworld",
		Port:   8080,
		Region: "eu-west-1",
		Repo:   "github.com/asim/helloworld",
	})
	fmt.Println(rsp, err)
}
