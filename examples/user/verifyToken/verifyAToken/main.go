package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Check whether the token attached to MagicLink is valid or not.
// Ideally, you need to call this endpoint from your http request
// handler that handles the endpoint which is specified in the
// SendMagicLink request.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.VerifyToken(&user.VerifyTokenRequest{
		Token: "EdsUiidouJJJLldjlloofUiorkojflsWWdld",
	})
	fmt.Println(rsp, err)
}
