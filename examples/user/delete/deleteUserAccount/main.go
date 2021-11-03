package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Delete an account by id
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Delete(&user.DeleteRequest{
		Id: "fdf34f34f34-f34f34-f43f43f34-f4f34f",
	})
	fmt.Println(rsp, err)
}
