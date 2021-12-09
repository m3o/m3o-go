package main

import (
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Vote to have the App api launched faster!
func main() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Vote(&app.VoteRequest{
		Message: "Launch it!",
	})
	fmt.Println(rsp, err)
}
