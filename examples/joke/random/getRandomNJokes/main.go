package main

import (
	"fmt"
	"os"

	"go.m3o.com/joke"
)

// Get a random joke
func main() {
	jokeService := joke.NewJokeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := jokeService.Random(&joke.RandomRequest{
		Count: 3,
	})
	fmt.Println(rsp, err)
}
