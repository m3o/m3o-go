package minecraft

import (
	"go.m3o.com/client"
)

type Minecraft interface {
	Ping(*PingRequest) (*PingResponse, error)
}

func NewMinecraftService(token string) *MinecraftService {
	return &MinecraftService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type MinecraftService struct {
	client *client.Client
}

// Ping a minecraft server
func (t *MinecraftService) Ping(request *PingRequest) (*PingResponse, error) {

	rsp := &PingResponse{}
	return rsp, t.client.Call("minecraft", "Ping", request, rsp)

}

type PingRequest struct {
	// address of the server
	Address string `json:"address,omitempty"`
}

type PingResponse struct {
	// Favicon in base64
	Favicon string `json:"favicon,omitempty"`
	// Latency (ms) between us and the server (EU)
	Latency int32 `json:"latency,omitempty"`
	// Max players ever
	MaxPlayers int32 `json:"max_players,omitempty"`
	// Message of the day
	Motd string `json:"motd,omitempty"`
	// Number of players online
	Players int32 `json:"players,omitempty"`
	// Protocol number of the server
	Protocol int32 `json:"protocol,omitempty"`
	// List of connected players
	Sample []PlayerSample `json:"sample,omitempty"`
	// Version of the server
	Version string `json:"version,omitempty"`
}

type PlayerSample struct {
	// name of the player
	Name string `json:"name,omitempty"`
	// unique id of player
	Uuid string `json:"uuid,omitempty"`
}
