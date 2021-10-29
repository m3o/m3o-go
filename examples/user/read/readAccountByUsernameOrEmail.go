package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAccountByUsernameOrEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Username: "usrname-1",
	})
	fmt.Println(rsp, err)
}
