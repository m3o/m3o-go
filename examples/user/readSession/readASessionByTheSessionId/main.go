package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read a session by the session id. In the event it has expired or is not found and error is returned.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.ReadSession(&user.ReadSessionRequest{
		SessionId: "sds34s34s34-s34s34-s43s43s34-s4s34s",
	})
	fmt.Println(rsp, err)

}
