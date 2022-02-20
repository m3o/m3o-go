package main

import (
	"fmt"
	"os"

	"go.m3o.com/memegen"
)

//
func main() {
	memegenService := memegen.NewMemegenService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memegenService.Generate(&memegen.GenerateRequest{
		Id: "444501",
	})
	fmt.Println(rsp, err)
}
