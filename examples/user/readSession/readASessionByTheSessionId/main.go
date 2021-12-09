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
		SessionId: "df91a612-5b24-4634-99ff-240220ab8f55",
	})
	fmt.Println(rsp, err)
}
