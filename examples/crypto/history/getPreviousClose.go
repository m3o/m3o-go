package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/crypto"
)

// Returns the history for the previous close
func GetPreviousClose() {
	cryptoService := crypto.NewCryptoService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cryptoService.History(&crypto.HistoryRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}
