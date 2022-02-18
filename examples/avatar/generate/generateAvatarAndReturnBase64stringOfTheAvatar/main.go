package main

import (
	"fmt"
	"os"

	"go.m3o.com/avatar"
)

// Generate an unique avatar
func main() {
	avatarService := avatar.NewAvatarService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := avatarService.Generate(&avatar.GenerateRequest{
		Format:   "jpeg",
		Gender:   "female",
		Upload:   false,
		Username: "",
	})
	fmt.Println(rsp, err)
}
