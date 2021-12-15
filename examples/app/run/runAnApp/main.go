package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Run an app from a source repo. Specify region etc.
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Run(&app.RunRequest{
		Branch: "master",
		Name:   "helloworld",
		Port:   8080,
		Region: "europe-west1",
		Repo:   "github.com/asim/helloworld",
	})
	fmt.Println(rsp, err)
}
