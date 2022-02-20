# Memegen

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/memegen/api](https://m3o.com/memegen/api).

Endpoints:

## Templates

List the available templates


[https://m3o.com/memegen/api#Templates](https://m3o.com/memegen/api#Templates)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/memegen"
)

// List the available templates
func MemeTemplates() {
	memegenService := memegen.NewMemegenService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memegenService.Templates(&memegen.TemplatesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Generate




[https://m3o.com/memegen/api#Generate](https://m3o.com/memegen/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/memegen"
)

// 
func GenerateAmeme() {
	memegenService := memegen.NewMemegenService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memegenService.Generate(&memegen.GenerateRequest{
		Id: "444501",

	})
	fmt.Println(rsp, err)
	
}
```
