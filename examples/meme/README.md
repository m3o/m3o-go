# Meme

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/meme/api](https://m3o.com/meme/api).

Endpoints:

## Templates

List the available templates


[https://m3o.com/meme/api#Templates](https://m3o.com/meme/api#Templates)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/meme"
)

// List the available templates
func MemeTemplates() {
	memeService := meme.NewMemeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memeService.Templates(&meme.TemplatesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Generate




[https://m3o.com/meme/api#Generate](https://m3o.com/meme/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/meme"
)

// 
func GenerateAmeme() {
	memeService := meme.NewMemeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memeService.Generate(&meme.GenerateRequest{
		Id: "444501",

	})
	fmt.Println(rsp, err)
	
}
```
