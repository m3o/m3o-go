package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/currency"
)

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
func ConvertUsdToGbp() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.Convert(&currency.ConvertRequest{
		From: "USD",
		To:   "GBP",
	})
	fmt.Println(rsp, err)
}
