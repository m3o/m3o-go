# Ai

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/ai/api](https://m3o.com/ai/api).

Endpoints:

## Call

Make a request to the AI


[https://m3o.com/ai/api#Call](https://m3o.com/ai/api#Call)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ai"
)

// Make a request to the AI
func CallTheAi() {
	aiService := ai.NewAiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := aiService.Call(&ai.CallRequest{
		Text: "who is leonardo davinci",

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
