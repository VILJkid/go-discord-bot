package main

import (
	"fmt"

	"github.com/VILJkid/go-discord-bot/app"
)

func main() {
	err := app.RunBot()
	if err != nil {
		fmt.Printf("Error starting the bot: %s\n", err)
	}
}
