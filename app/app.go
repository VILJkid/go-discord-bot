package app

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/VILJkid/go-discord-bot/pkg/bot"
	"github.com/VILJkid/go-discord-bot/pkg/events"
	"github.com/VILJkid/go-discord-bot/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

func preFlightCheck(s *discordgo.Session) (err error) {
	botUser, err := s.User("@me")
	if err != nil {
		slog.Error("Unable to fetch bot details")
		return
	}

	header := []string{"User ID", "Username", "Discriminator"}
	data := []string{botUser.ID, botUser.Username, botUser.Discriminator}

	utils.PrintInTable(header, data)

	fmt.Println("Press \"c\" to continue or any other key to exit...")
	var input string
	fmt.Scanln(&input)

	if input != "c" && input != "C" {
		err = errors.New("manual interruption by user")
		slog.Warn("Exiting...")
		return
	}
	return
}

func RunBot() (err error) {
	config, err := utils.GetConfigs()
	if err != nil {
		return
	}
	// Initialize DiscordGo session
	if len(config.BotToken) == 0 {
		err = errors.New("bot token not set or provided")
		return
	}
	s, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		slog.Error("Unable to create a Discord session")
		return
	}

	if err = preFlightCheck(s); err != nil {
		return err
	}

	// Set up event handlers
	s.AddHandler(bot.OnMessageCreate)

	// Register the UserJoin event handler
	s.AddHandler(events.HandleUserJoin)

	s.AddHandler(bot.OnInteractionCreate)

	s.Identify.Intents = discordgo.IntentGuildMessages

	// Open the Discord connection
	slog.Info("Starting the Bot...")
	if err = s.Open(); err != nil {
		slog.Error("Unable to open the websocket connection")
		return
	}
	slog.Info("Bot is now running. Press \"CTRL+C\" to exit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	slog.Warn("Shutting down the bot...")
	s.Close()
	slog.Info("Bot is now stopped")
	return
}
