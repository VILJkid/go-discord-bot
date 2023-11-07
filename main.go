package main

import (
	"log/slog"

	"github.com/VILJkid/go-discord-bot/app"
)

func main() {
	err := app.RunBot()
	if err != nil {
		slog.Error("Failed starting the bot:", "reason", err)
	}
}
