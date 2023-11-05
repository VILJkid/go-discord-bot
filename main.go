package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"
)

var (
	Token string
)

func init() {
	pflag.StringVarP(&Token, "token", "t", "", "Authenticate using Discord bot token")
	pflag.Parse()
	pflag.Usage = Help
}

func Help() {
	fmt.Println("\nUsage: bot [options]")
	fmt.Println("Options:")
	fmt.Println()
	data := [][]string{}
	pflag.VisitAll(func(f *pflag.Flag) {
		v := f.DefValue
		if len(v) == 0 {
			v = "Not set"
		}

		data = append(data, []string{
			"--" + f.Name,
			"-" + f.Shorthand,
			f.Value.Type(),
			v,
			f.Usage,
		})
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flag", "Shorthand", "Type", "Default Value", "Description"})
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	commandPrefix := "!"

	if !strings.HasPrefix(m.Content, commandPrefix) {
		return
	}

	parts := strings.Fields(m.Content)
	if len(parts) < 1 {
		return
	}

	command := strings.TrimPrefix(parts[0], commandPrefix)
	switch command {
	case "ping":
		slog.Info("Command recieved:", "command", "ping")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		slog.Info("Response sent:", "message", "Pong!")
	case "pong":
		slog.Info("Command recieved:", "command", "pong")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		slog.Info("Response sent:", "message", "Ping!")
	}
}

func Setup() (dg *discordgo.Session, err error) {
	if len(Token) == 0 {
		err = errors.New("bot token not set or provided")
		return
	}
	dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		slog.Error("Unable to create a Discord session:")
		return
	}
	

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentGuildMessages
	return
}

func Launch(dg *discordgo.Session) (err error) {
	slog.Info("Starting the Bot...")
	if err = dg.Open(); err != nil {
		slog.Error("Unable to open the websocket connection")
		return
	}
	slog.Info("Bot is now running. Press \"CTRL+C\" to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	slog.Warn("Shutting down the bot...")
	dg.Close()
	slog.Info("Bot is now stopped.")
	return
}

func main() {
	dg, err := Setup()
	if err != nil {
		slog.Error("Unable to setup the bot:", "reason", err)
		pflag.Usage()
		return
	}
	if err = Launch(dg); err != nil {
		slog.Error("Unable to launch the bot:", "reason", err)
		pflag.Usage()
		return
	}
}
