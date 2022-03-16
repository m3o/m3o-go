package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.Delete(&user.DeleteRequest{
		Id: "8b98acbe-0b6a-4d66-a414-5ffbf666786f",
	})
	fmt.Println(rsp, err)
}
