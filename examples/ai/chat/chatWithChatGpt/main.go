package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ai"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ai.Chat(&ai.ChatRequest{
		Context: []ai.Context{
			ai.Context{
				Prompt: "who is leonardo davinci",
				Reply:  "Leonardo da Vinci was an Italian polymath during the Renaissance period. He was born in 1452 in Vinci, Italy, and is renowned for his contributions to various fields such as science, engineering, art, and anatomy. Da Vinci's most famous works include the Mona Lisa and The Last Supper. He is often considered one of the greatest artists and thinkers of all time.",
			}},
		Model:  "gpt-3.5-turbo",
		Prompt: "when did he die?",
	})
	fmt.Println(rsp, err)
}
