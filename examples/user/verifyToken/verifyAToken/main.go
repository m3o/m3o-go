package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

// Check whether the token attached to MagicLink is valid or not.
// Ideally, you need to call this endpoint from your http request
// handler that handles the endpoint which is specified in the
// SendMagicLink request.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.VerifyToken(&user.VerifyTokenRequest{
		Token: "EdsUiidouJJJLldjlloofUiorkojflsWWdld",
	})
	fmt.Println(rsp, err)
}
