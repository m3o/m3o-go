# Function

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Function/api](https://m3o.com/Function/api).

Endpoints:

## List

List all the deployed functions


[https://m3o.com/function/api#List](https://m3o.com/function/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// List all the deployed functions
func ListFunctions() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.List(&function.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a function by name


[https://m3o.com/function/api#Delete](https://m3o.com/function/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Delete a function by name
func DeleteAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Delete(&function.DeleteRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Describe

Get the info for a deployed function


[https://m3o.com/function/api#Describe](https://m3o.com/function/api#Describe)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Get the info for a deployed function
func DescribeFunctionStatus() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Describe(&function.DescribeRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Deploy

Deploy a group of functions


[https://m3o.com/function/api#Deploy](https://m3o.com/function/api#Deploy)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Deploy a group of functions
func DeployAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Deploy(&function.DeployRequest{
		Entrypoint: "helloworld",
Name: "helloworld",
Repo: "github.com/m3o/nodejs-function-example",
Runtime: "nodejs14",

	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update a function


[https://m3o.com/function/api#Update](https://m3o.com/function/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Update a function
func UpdateAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Update(&function.UpdateRequest{
		Entrypoint: "helloworld",
Name: "helloworld",
Repo: "github.com/m3o/nodejs-function-example",
Runtime: "nodejs14",

	})
	fmt.Println(rsp, err)
	
}
```
## Call

Call a function by name


[https://m3o.com/function/api#Call](https://m3o.com/function/api#Call)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Call a function by name
func CallAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Call(&function.CallRequest{
		Name: "helloworld",
Request: map[string]interface{}{
},

	})
	fmt.Println(rsp, err)
	
}
```
