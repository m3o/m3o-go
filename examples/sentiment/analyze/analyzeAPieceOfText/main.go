package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/sentiment"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Sentiment.Analyze(&sentiment.AnalyzeRequest{
		Text: "this is amazing",
	})
	fmt.Println(rsp, err)
}
