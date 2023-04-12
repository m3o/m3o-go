# Cron

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/cron/api](https://m3o.com/cron/api).

Endpoints:

## Delete

Delete a cron job


[https://m3o.com/cron/api#Delete](https://m3o.com/cron/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cron"
)

// Delete a cron job
func DeleteAjob() {
	cronService := cron.NewCronService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cronService.Delete(&cron.DeleteRequest{
		Id: "0c8cf9f7-3a61-4e91-b249-00a970044c95",

	})
	fmt.Println(rsp, err)
	
}
```
## Jobs

List all cron jobs


[https://m3o.com/cron/api#Jobs](https://m3o.com/cron/api#Jobs)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cron"
)

// List all cron jobs
func ListJobs() {
	cronService := cron.NewCronService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cronService.Jobs(&cron.JobsRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Schedule

Schedule a cron job


[https://m3o.com/cron/api#Schedule](https://m3o.com/cron/api#Schedule)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cron"
)

// Schedule a cron job
func ScheduleAjob() {
	cronService := cron.NewCronService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cronService.Schedule(&cron.ScheduleRequest{
		Callback: "https://google.com",
Description: "testing",
Interval: "* * * * *",
Name: "test",

	})
	fmt.Println(rsp, err)
	
}
```
