package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/holidays"
)

// Get the list of countries that are supported by this API
func ListCountries() {
	holidaysService := holidays.NewHolidaysService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := holidaysService.Countries(&holidays.CountriesRequest{})
	fmt.Println(rsp, err)
}
