# Ai

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/ai/api](https://m3o.com/ai/api).

Endpoints:

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
		Model: "gpt-3.5-turbo",
Prompt: "who is leonardo davinci",

	})
	fmt.Println(rsp, err)
	
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
