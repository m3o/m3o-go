package main

import (
	"fmt"
	"os"

	"go.m3o.com/meme"
)

//
func main() {
	memeService := meme.NewMemeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memeService.Generate(&meme.GenerateRequest{
		Id: "444501",
	})
	fmt.Println(rsp, err)
}
