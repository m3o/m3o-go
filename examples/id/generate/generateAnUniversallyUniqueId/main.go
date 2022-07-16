package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/id"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Id.Generate(&id.GenerateRequest{
		Type: "uuid",
	})
	fmt.Println(rsp, err)
}
