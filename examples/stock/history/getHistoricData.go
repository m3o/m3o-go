package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/stock"
)

// Get the historic open-close for a given day
func GetHistoricData() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.History(&stock.HistoryRequest{
		Date:  "2020-10-01",
		Stock: "AAPL",
	})
	fmt.Println(rsp, err)
}
