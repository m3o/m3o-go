# Secret

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/secret/api](https://m3o.com/secret/api).

Endpoints:

## List

List all the available secrets


[https://m3o.com/secret/api#List](https://m3o.com/secret/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/secret"
)

// List all the available secrets
func ListAllSecrets() {
	secretService := secret.NewSecretService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := secretService.List(&secret.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Set

Set a secret. Overwrites any existing value already set.


[https://m3o.com/secret/api#Set](https://m3o.com/secret/api#Set)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/secret"
)

// Set a secret. Overwrites any existing value already set.
func SetAvalue() {
	secretService := secret.NewSecretService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := secretService.Set(&secret.SetRequest{
		Key: "foo",
Value: "bar",

	})
	fmt.Println(rsp, err)
	
}
```
## Get

Get a secret by key.


[https://m3o.com/secret/api#Get](https://m3o.com/secret/api#Get)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/secret"
)

// Get a secret by key.
func GetAvalue() {
	secretService := secret.NewSecretService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := secretService.Get(&secret.GetRequest{
		Key: "foo",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a secret. If key not found a success response is returned.


[https://m3o.com/secret/api#Delete](https://m3o.com/secret/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/secret"
)

// Delete a secret. If key not found a success response is returned.
func DeleteAvalue() {
	secretService := secret.NewSecretService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := secretService.Delete(&secret.DeleteRequest{
		Key: "foo",

	})
	fmt.Println(rsp, err)
	
}
```
