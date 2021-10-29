package example

import (
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Returns the data for the previous close
func GetPreviousClose() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.History(&forex.HistoryRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
