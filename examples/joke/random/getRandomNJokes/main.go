package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/joke"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Joke.Random(&joke.RandomRequest{
		Count: 3,
	})
	fmt.Println(rsp, err)
}
