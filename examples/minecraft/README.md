# Minecraft

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/minecraft/api](https://m3o.com/minecraft/api).

Endpoints:

## Ping

Ping a minecraft server


[https://m3o.com/minecraft/api#Ping](https://m3o.com/minecraft/api#Ping)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/minecraft"
)

// Ping a minecraft server
func PingAminecraftServer() {
	minecraftService := minecraft.NewMinecraftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := minecraftService.Ping(&minecraft.PingRequest{
		Address: "funcraft.net",

	})
	fmt.Println(rsp, err)
	
}
```
