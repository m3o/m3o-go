# Password

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/password/api](https://m3o.com/password/api).

Endpoints:

## Generate

Generate a strong random password


[https://m3o.com/password/api#Generate](https://m3o.com/password/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/password"
)

// Generate a strong random password
func GeneratePassword() {
	passwordService := password.NewPasswordService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := passwordService.Generate(&password.GenerateRequest{
		Length: 16,

	})
	fmt.Println(rsp, err)
	
}
```
