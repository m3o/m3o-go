package example

import (
	"fmt"
	"os"

	"go.m3o.com/vehicle"
)

// Lookup a UK vehicle by it's registration number
func LookupVehicle() {
	vehicleService := vehicle.NewVehicleService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := vehicleService.Lookup(&vehicle.LookupRequest{
		Registration: "LC60OTA",
	})
	fmt.Println(rsp, err)
}
