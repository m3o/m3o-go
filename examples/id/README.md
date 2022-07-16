# Id

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/id/api](https://m3o.com/id/api).

Endpoints:

## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAnUniversallyUniqueId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "uuid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAnUniqueLexicographicallyId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "ulid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAnUniqueKsortableId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "ksuid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAwebOrientedGloballyUniqueId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "xid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAurlFriendlyUniqueId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "nanoid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAshortId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "shortid",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAsnowflakeId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "snowflake",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a unique ID. Defaults to uuid.


[https://m3o.com/id/api#Generate](https://m3o.com/id/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAbigflakeId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "bigflake",

	})
	fmt.Println(rsp, err)
	
}
```
## Types

List the types of IDs available.


[https://m3o.com/id/api#Types](https://m3o.com/id/api#Types)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/id"
)

// List the types of IDs available.
func ListTheTypesOfIdsAvailable() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Types(&id.TypesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
