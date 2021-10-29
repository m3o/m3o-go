package example

import (
	"fmt"
	"os"

	"go.m3o.com/holidays"
)

// List the holiday dates for a given country and year
func GetHolidays() {
	holidaysService := holidays.NewHolidaysService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := holidaysService.List(&holidays.ListRequest{
		Year: 2022,
	})
	fmt.Println(rsp, err)
}
