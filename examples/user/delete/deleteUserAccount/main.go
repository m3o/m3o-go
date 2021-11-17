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
		Id: "8b98acbe-0b6a-4d66-a414-5ffbf666786f",
	})
	fmt.Println(rsp, err)

}
