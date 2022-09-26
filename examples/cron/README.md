# Cron

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/cron/api](https://m3o.com/cron/api).

Endpoints:

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
func ListJobs() {
	cronService := cron.NewCronService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cronService.Schedule(&cron.ScheduleRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
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
