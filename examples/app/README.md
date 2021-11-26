# App

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/App/api](https://m3o.com/App/api).

Endpoints:

## Vote

Vote to have the App api launched faster!


[https://m3o.com/app/api#Vote](https://m3o.com/app/api#Vote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Vote to have the App api launched faster!
func VoteForTheApi() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Vote(&app.VoteRequest{
		Message: "Launch it!",

	})
	fmt.Println(rsp, err)
	
}
```
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
