package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/id"
)

// Generate a unique ID. Defaults to uuid.
func GenerateAsnowflakeId() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Generate(&id.GenerateRequest{
		Type: "snowflake",
	})
	fmt.Println(rsp, err)
}
