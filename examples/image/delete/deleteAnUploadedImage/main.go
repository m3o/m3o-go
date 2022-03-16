package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/image"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Image.Delete(&image.DeleteRequest{
		Url: "https://cdn.m3ocontent.com/micro/images/micro/41e23b39-48dd-42b6-9738-79a313414bb8/cat.png",
	})
	fmt.Println(rsp, err)
}
