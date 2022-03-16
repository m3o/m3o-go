package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Resolve an app by id to its raw backend endpoint
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Resolve(&app.ResolveRequest{
		Id: "helloworld",
	})
	fmt.Println(rsp, err)
}
