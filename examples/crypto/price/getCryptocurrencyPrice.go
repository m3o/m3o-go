package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/crypto"
)

// Get the last price for a given crypto ticker
func GetCryptocurrencyPrice() {
	cryptoService := crypto.NewCryptoService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cryptoService.Price(&crypto.PriceRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}
