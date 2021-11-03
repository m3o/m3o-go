package main

import (
	"fmt"
	"os"

	"go.m3o.com/sentiment"
)

// Analyze and score a piece of text
func main() {
	sentimentService := sentiment.NewSentimentService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := sentimentService.Analyze(&sentiment.AnalyzeRequest{
		Text: "this is amazing",
	})
	fmt.Println(rsp, err)

}
