package main

import (
	"fmt"
	"os"

	"go.m3o.com/minecraft"
)

// Ping a minecraft server
func main() {
	minecraftService := minecraft.NewMinecraftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := minecraftService.Ping(&minecraft.PingRequest{
		Address: "funcraft.net",
	})
	fmt.Println(rsp, err)
}
