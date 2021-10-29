package example

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Update the account username or email
func UpdateAnAccount() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Update(&user.UpdateRequest{
		Email: "joeotheremail@example.com",
		Id:    "usrid-1",
	})
	fmt.Println(rsp, err)
}
