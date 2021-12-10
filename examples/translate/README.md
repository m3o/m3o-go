# Translate

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Translate/api](https://m3o.com/Translate/api).

Endpoints:

## Text

Basic text translation


[https://m3o.com/translate/api#Text](https://m3o.com/translate/api#Text)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/translate"
)

// Basic text translation
func TranslateText() {
	translateService := translate.NewTranslateService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := translateService.Text(&translate.TextRequest{
		Content: "hello",
Format: "text",
Model: "nmt",
Source: "en",
Target: "fr",

	})
	fmt.Println(rsp, err)
	
}
```
