package main

import (
	"fmt"
	"os"

	"go.m3o.com/answer"
)

// Ask a question and receive an instant answer
func main() {
	answerService := answer.NewAnswerService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := answerService.Question(&answer.QuestionRequest{
		Query: "microsoft",
	})
	fmt.Println(rsp, err)
}
