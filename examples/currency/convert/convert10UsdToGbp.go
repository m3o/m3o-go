package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/currency"
)

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
func Convert10usdToGbp() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.Convert(&currency.ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "GBP",
	})
	fmt.Println(rsp, err)
}
