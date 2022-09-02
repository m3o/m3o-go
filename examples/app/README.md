# App

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/app/api](https://m3o.com/app/api).

Endpoints:

## Logs

Get the logs for an app


[https://m3o.com/app/api#Logs](https://m3o.com/app/api#Logs)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Get the logs for an app
func RetrieveBuildLogsForAnApp() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Logs(&app.LogsRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Reserve

Reserve app names


[https://m3o.com/app/api#Reserve](https://m3o.com/app/api#Reserve)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Reserve app names
func ReserveAppName() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Reserve(&app.ReserveRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Regions

Return the support regions


[https://m3o.com/app/api#Regions](https://m3o.com/app/api#Regions)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Return the support regions
func ListRegions() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Regions(&app.RegionsRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Status

Get the status of an app


[https://m3o.com/app/api#Status](https://m3o.com/app/api#Status)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Get the status of an app
func GetTheStatusOfAnApp() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Status(&app.StatusRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Resolve

Resolve an app by id to its raw backend endpoint


[https://m3o.com/app/api#Resolve](https://m3o.com/app/api#Resolve)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Resolve an app by id to its raw backend endpoint
func ResolveAppById() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Resolve(&app.ResolveRequest{
		Id: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update the app. The latest source code will be downloaded, built and deployed.


[https://m3o.com/app/api#Update](https://m3o.com/app/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Update the app. The latest source code will be downloaded, built and deployed.
func UpdateAnApp() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Update(&app.UpdateRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List all the apps


[https://m3o.com/app/api#List](https://m3o.com/app/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// List all the apps
func ListTheApps() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.List(&app.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Run

Run an app from source


[https://m3o.com/app/api#Run](https://m3o.com/app/api#Run)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Run an app from source
func RunAnApp() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Run(&app.RunRequest{
		Branch: "master",
Name: "helloworld",
Port: 8080,
Region: "europe-west1",
Repo: "github.com/asim/helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete an app


[https://m3o.com/app/api#Delete](https://m3o.com/app/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/app"
)

// Delete an app
func DeleteAnApp() {
	appService := app.NewAppService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := appService.Delete(&app.DeleteRequest{
		Name: "helloworld",

	})
	fmt.Println(rsp, err)
	
}
```
