package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/answer"
)

// Ask a question and receive an instant answer
func AskAquestion() {
	answerService := answer.NewAnswerService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := answerService.Question(&answer.QuestionRequest{
		Query: "microsoft",
	})
	fmt.Println(rsp, err)
}
