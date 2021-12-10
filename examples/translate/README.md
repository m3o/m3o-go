# Translate

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Translate/api](https://m3o.com/Translate/api).

Endpoints:

## Text

TextRequest is the basic edition request


[https://m3o.com/translate/api#Text](https://m3o.com/translate/api#Text)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/translate"
)

// TextRequest is the basic edition request
func TranslateString() {
	translateService := translate.NewTranslateService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := translateService.Text(&translate.TextRequest{
		Format: "text",
Model: "nmt",
Source: "en",
Target: "zh",

	})
	fmt.Println(rsp, err)
	
}
```