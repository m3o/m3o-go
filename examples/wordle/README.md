# Wordle

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/wordle/api](https://m3o.com/wordle/api).

Endpoints:

## Next

When does the next game start


[https://m3o.com/wordle/api#Next](https://m3o.com/wordle/api#Next)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wordle"
)

// When does the next game start
func NextGame() {
	wordleService := wordle.NewWordleService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := wordleService.Next(&wordle.NextRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Guess

Make a guess


[https://m3o.com/wordle/api#Guess](https://m3o.com/wordle/api#Guess)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wordle"
)

// Make a guess
func MakeAguess() {
	wordleService := wordle.NewWordleService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := wordleService.Guess(&wordle.GuessRequest{
		Word: "noise",

	})
	fmt.Println(rsp, err)
	
}
```
