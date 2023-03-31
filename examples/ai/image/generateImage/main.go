package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ai"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ai.Image(&ai.ImageRequest{
		Text: "a cat on a ball",
	})
	fmt.Println(rsp, err)
}
