# Email

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/email/api](https://m3o.com/email/api).

Endpoints:

## Send

Send an email by passing in from, to, subject, and a text or html body


[https://m3o.com/email/api#Send](https://m3o.com/email/api#Send)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Send an email by passing in from, to, subject, and a text or html body
func SendEmail() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Send(&email.SendRequest{
		From: "Awesome Dot Com",
Subject: "Email verification",
TextBody: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,

	})
	fmt.Println(rsp, err)
	
}
```
## Parse

Parse an RFC5322 address e.g "Joe Blogs <joe@example.com>"


[https://m3o.com/email/api#Parse](https://m3o.com/email/api#Parse)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Parse an RFC5322 address e.g "Joe Blogs <joe@example.com>"
func ParseEmail() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Parse(&email.ParseRequest{
		Address: "Joe Blogs <joe@example.com>",

	})
	fmt.Println(rsp, err)
	
}
```
## Validate

Validate an email address format


[https://m3o.com/email/api#Validate](https://m3o.com/email/api#Validate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Validate an email address format
func ValidateEmail() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Validate(&email.ValidateRequest{
		Address: "joe@example.com",

	})
	fmt.Println(rsp, err)
	
}
```
