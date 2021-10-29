package example

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAccountByEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Email: "joe@example.com",
	})
	fmt.Println(rsp, err)
}
