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
	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
)

func botDetails(s *discordgo.Session) (err error) {
	botUser, err := s.User("@me")
	if err != nil {
		slog.Error("Unable to fetch bot details")
		return
	}

	header := []string{"User ID", "Username", "Discriminator"}
	data := []string{botUser.ID, botUser.Username, botUser.Discriminator}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(true)
	table.Append(data)
	table.Render()
	return
}

func RunBot() (err error) {
	Token := ""
	// Initialize DiscordGo session
	if len(Token) == 0 {
		err = errors.New("bot token not set or provided")
		return
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		slog.Error("Unable to create a Discord session")
		return
	}

	if err = botDetails(dg); err != nil {
		return err
	}

	fmt.Println("Press \"c\" to continue or any other key to exit...")
	var input string
	fmt.Scanln(&input)

	if input != "c" && input != "C" {
		err = errors.New("manual interruption by user")
		slog.Warn("Exiting...")
		return
	}

	// Set up event handlers
	dg.AddHandler(bot.OnMessageCreate)

	// Register the UserJoin event handler
	dg.AddHandler(events.HandleUserJoin)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	// Open the Discord connection
	slog.Info("Starting the Bot...")
	if err = dg.Open(); err != nil {
		slog.Error("Unable to open the websocket connection")
		return
	}
	slog.Info("Bot is now running. Press \"CTRL+C\" to exit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	slog.Warn("Shutting down the bot...")
	dg.Close()
	slog.Info("Bot is now stopped")
	return
}
