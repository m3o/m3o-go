package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ai"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ai.Chat(&ai.ChatRequest{
		Model:  "gpt-3.5-turbo",
		Prompt: "who is leonardo davinci",
	})
	fmt.Println(rsp, err)
}
