package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cron"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cron.Delete(&cron.DeleteRequest{
		Id: "0c8cf9f7-3a61-4e91-b249-00a970044c95",
	})
	fmt.Println(rsp, err)
}
