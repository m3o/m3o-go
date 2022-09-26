package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cron"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cron.Schedule(&cron.ScheduleRequest{})
	fmt.Println(rsp, err)
}
