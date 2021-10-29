package example

import (
	"fmt"
	"os"

	"go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAshortId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "shortid",
	})
	fmt.Println(rsp, err)
}
