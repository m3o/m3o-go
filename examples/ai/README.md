# Ai

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/ai/api](https://m3o.com/ai/api).

Endpoints:

## Edit

Edit or edit prompt/code


[https://m3o.com/ai/api#Edit](https://m3o.com/ai/api#Edit)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Edit or edit prompt/code
func EditText() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Edit(&ai.EditRequest{
		Text: "What day of the wek is it?",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate an image from prompt


[https://m3o.com/ai/api#Generate](https://m3o.com/ai/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Generate an image from prompt
func GenerateImage() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Generate(&ai.GenerateRequest{
		Text: "a cat on a ball",

	})
	fmt.Println(rsp, err)
	
}
```
## Moderate

Moderate hate speech


[https://m3o.com/ai/api#Moderate](https://m3o.com/ai/api#Moderate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Moderate hate speech
func ModerateHateSpeech() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Moderate(&ai.ModerateRequest{
		Text: "I want to kill him",

	})
	fmt.Println(rsp, err)
	
}
```
## Chat

Make a request to ChatGPT


[https://m3o.com/ai/api#Chat](https://m3o.com/ai/api#Chat)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Make a request to ChatGPT
func ChatWithChatGpt() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Chat(&ai.ChatRequest{
		Context: []ai.Context{
ai.Context{
		Prompt: "who is leonardo davinci",
		Reply: "Leonardo da Vinci was an Italian polymath during the Renaissance period. He was born in 1452 in Vinci, Italy, and is renowned for his contributions to various fields such as science, engineering, art, and anatomy. Da Vinci's most famous works include the Mona Lisa and The Last Supper. He is often considered one of the greatest artists and thinkers of all time.",
}},
Model: "gpt-3.5-turbo",
Prompt: "when did he die?",

	})
	fmt.Println(rsp, err)
	
}
```
## Stream

Stream a response from chatgpt


[https://m3o.com/ai/api#Stream](https://m3o.com/ai/api#Stream)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Stream a response from chatgpt
func StreamResponseFromChatGpt() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := aiService.Stream(&ai.StreamRequest{
		Model: "gpt-3.5-turbo",
Prompt: "write helloworld in node.js",

	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
			rsp, err := stream.Recv()
			if err != nil {
					fmt.Println(err)
					return
			}

			fmt.Println(rsp)
	}
}
```
## Complete

Make a request to the AI


[https://m3o.com/ai/api#Complete](https://m3o.com/ai/api#Complete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Make a request to the AI
func CompleteTheText() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Complete(&ai.CompleteRequest{
		Text: "who is leonardo davinci",

	})
	fmt.Println(rsp, err)
	
}
```
