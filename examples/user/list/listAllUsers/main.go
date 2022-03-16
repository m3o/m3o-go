package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.List(&user.ListRequest{
		Limit:  100,
		Offset: 0,
	})
	fmt.Println(rsp, err)
}
