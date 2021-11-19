package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// List all users. Returns a paged list of results
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.List(&user.ListRequest{
		Limit:  100,
		Offset: 0,
	})
	fmt.Println(rsp, err)

}
