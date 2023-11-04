package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
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

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func Setup() (dg *discordgo.Session, err error) {
	if len(Token) == 0 {
		err = errors.New("token not set or provided")
		return
	}
	dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		slog.Error("unable to create a Discord session:")
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentGuildMessages
	return
}

func Launch(dg *discordgo.Session) (err error) {
	if err = dg.Open(); err != nil {
		slog.Error("unable to open the websocket connection")
		return
	}
	slog.Info("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
	return
}

func main() {
	dg, err := Setup()
	if err != nil {
		slog.Error("unable to setup the bot:", "reason", err)
		Help()
		return
	}
	if err = Launch(dg); err != nil {
		slog.Error("unable to launch the bot:", "reason", err)
		Help()
		return
	}
}
