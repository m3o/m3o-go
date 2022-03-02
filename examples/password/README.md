# Password

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/password/api](https://m3o.com/password/api).

Endpoints:

## Generate

Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them


[https://m3o.com/password/api#Generate](https://m3o.com/password/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/password"
)

// Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them
func GeneratePassword() {
	passwordService := password.NewPasswordService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := passwordService.Generate(&password.GenerateRequest{
		Length: 16,

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them


[https://m3o.com/password/api#Generate](https://m3o.com/password/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/password"
)

// Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them
func GeneratePasswordWithoutSpecialCharacters() {
	passwordService := password.NewPasswordService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := passwordService.Generate(&password.GenerateRequest{
		Length: 16,
Lowercase: true,
Numbers: true,
Special: false,
Uppercase: true,

	})
	fmt.Println(rsp, err)
	
}
```
