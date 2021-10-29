package example

import (
	"fmt"
	"os"

	"go.m3o.com/evchargers"
)

// Retrieve reference data as used by this API and in conjunction with the Search endpoint
func GetReferenceData() {
	evchargersService := evchargers.NewEvchargersService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := evchargersService.ReferenceData(&evchargers.ReferenceDataRequest{})
	fmt.Println(rsp, err)
}
