# Spam

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/spam/api](https://m3o.com/spam/api).

Endpoints:

## Classify

Check whether an email is likely to be spam based on its attributes


[https://m3o.com/spam/api#Classify](https://m3o.com/spam/api#Classify)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/spam"
)

// Check whether an email is likely to be spam based on its attributes
func ClassifyAnEmail() {
	spamService := spam.NewSpamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spamService.Classify(&spam.ClassifyRequest{
		From: "noreply@m3o.com",
Subject: "Welcome",
To: "hello@example.com",

	})
	fmt.Println(rsp, err)
	
}
```
## Classify

Check whether an email is likely to be spam based on its attributes


[https://m3o.com/spam/api#Classify](https://m3o.com/spam/api#Classify)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/spam"
)

// Check whether an email is likely to be spam based on its attributes
func ClassifyAnEmailUsingTheRawData() {
	spamService := spam.NewSpamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spamService.Classify(&spam.ClassifyRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
