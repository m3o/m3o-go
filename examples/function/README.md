# Function

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/function/api](https://m3o.com/function/api).

Endpoints:

## Update

Update a function. Downloads the source, builds and redeploys


[https://m3o.com/function/api#Update](https://m3o.com/function/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Update a function. Downloads the source, builds and redeploys
func UpdateAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Update(&function.UpdateRequest{
		Name: "helloworld",

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
	"name": "Alice",
},

	})
	fmt.Println(rsp, err)
	
}
```
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
## Proxy

Return the backend url for proxying


[https://m3o.com/function/api#Proxy](https://m3o.com/function/api#Proxy)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Return the backend url for proxying
func ProxyUrl() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Proxy(&function.ProxyRequest{
		Id: "helloworld",

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
		Branch: "main",
Entrypoint: "Helloworld",
Name: "helloworld",
Region: "europe-west1",
Repo: "https://github.com/m3o/m3o",
Runtime: "go116",
Subfolder: "examples/go-function",

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
## Regions

Return a list of supported regions


[https://m3o.com/function/api#Regions](https://m3o.com/function/api#Regions)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Return a list of supported regions
func ListRegions() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Regions(&function.RegionsRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Reserve

Reserve function names and resources beyond free quota


[https://m3o.com/function/api#Reserve](https://m3o.com/function/api#Reserve)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Reserve function names and resources beyond free quota
func ReserveAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Reserve(&function.ReserveRequest{
		Name: "helloworld",

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
