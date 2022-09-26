package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cron"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cron.Jobs(&cron.JobsRequest{})
	fmt.Println(rsp, err)
}
