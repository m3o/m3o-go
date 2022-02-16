# Otp

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/otp/api](https://m3o.com/otp/api).

Endpoints:

## Generate

Generate an OTP (one time pass) code


[https://m3o.com/otp/api#Generate](https://m3o.com/otp/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/otp"
)

// Generate an OTP (one time pass) code
func GenerateOtp() {
	otpService := otp.NewOtpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := otpService.Generate(&otp.GenerateRequest{
		Id: "asim@example.com",

	})
	fmt.Println(rsp, err)
	
}
```
## Validate

Validate the OTP code


[https://m3o.com/otp/api#Validate](https://m3o.com/otp/api#Validate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/otp"
)

// Validate the OTP code
func ValidateOtp() {
	otpService := otp.NewOtpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := otpService.Validate(&otp.ValidateRequest{
		Code: "656211",
Id: "asim@example.com",

	})
	fmt.Println(rsp, err)
	
}
```
