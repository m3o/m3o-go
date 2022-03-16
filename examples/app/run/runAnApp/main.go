package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Run an app from source
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Run(&app.RunRequest{
		Branch: "master",
		Name:   "helloworld",
		Port:   8080,
		Region: "europe-west1",
		Repo:   "github.com/asim/helloworld",
	})
	fmt.Println(rsp, err)
}
