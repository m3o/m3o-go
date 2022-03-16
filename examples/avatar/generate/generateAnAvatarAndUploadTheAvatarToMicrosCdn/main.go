package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/avatar"
)

// Generate an unique avatar
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Avatar.Generate(&avatar.GenerateRequest{
		Format:   "png",
		Gender:   "female",
		Upload:   true,
		Username: "",
	})
	fmt.Println(rsp, err)
}
