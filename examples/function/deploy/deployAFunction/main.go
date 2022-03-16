package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Deploy(&function.DeployRequest{
		Branch:     "main",
		Entrypoint: "Helloworld",
		Name:       "helloworld",
		Region:     "europe-west1",
		Repo:       "https://github.com/m3o/m3o",
		Runtime:    "go116",
		Subfolder:  "examples/go-function",
	})
	fmt.Println(rsp, err)
}
