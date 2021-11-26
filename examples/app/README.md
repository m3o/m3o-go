# App

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/App/api](https://m3o.com/App/api).

Endpoints:

## Reserve

Reserve your app name


[https://m3o.com/app/api#Reserve](https://m3o.com/app/api#Reserve)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Reserve your app name
func ReserveAppName() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Reserve(&app.ReserveRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
