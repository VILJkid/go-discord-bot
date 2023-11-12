package bot

import (
	"context"
	"log/slog"
	"strings"

	"github.com/VILJkid/go-discord-bot/pkg/commands"
	"github.com/VILJkid/go-discord-bot/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

type Command interface {
	SetCommandConfig(context.Context, *discordgo.Session, *discordgo.MessageCreate, []string) error
	ValidateCommand() error
	ExecuteCommand() error
}

func ErrorMessageEmbed(description string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: description,
		Color:       16711680,
	}
}

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate, c string, a []string) {
	ctx := context.Background()
	ctx = utils.SetContextRequestID(ctx, uuid.NewString())

	slog.InfoContext(ctx, "Request started:", utils.ConstContextRequestID, utils.GetContextRequestID(ctx))
	defer slog.InfoContext(ctx, "Request finished:", utils.ConstContextRequestID, utils.GetContextRequestID(ctx))

	var command Command
	c = strings.ToLower(c)
	switch c {
	case utils.CommandPing:
		command = new(commands.Ping)
	case utils.CommandPong:
		command = new(commands.Pong)
	case utils.CommandCreateChannel:
		command = new(commands.CreateChannel)
	case utils.CommandCreateChannelButton:
		command = new(commands.CreateChannelButton)
	default:
		embed := ErrorMessageEmbed("Command not recognized: " + c)
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		slog.Error("Command not recognized:", "command", c)
		return
	}

	slog.InfoContext(ctx, "Command recieved:", "command", c)

	if err := command.SetCommandConfig(ctx, s, m, a); err != nil {
		slog.ErrorContext(ctx, "Unable to set the command config:", "command", c, "reason", err)
		return
	}
	if err := command.ValidateCommand(); err != nil {
		slog.ErrorContext(ctx, "Unable to validate the command:", "command", c, "reason", err)
		return
	}
	if err := command.ExecuteCommand(); err != nil {
		slog.ErrorContext(ctx, "Unable to execute the command:", "command", c, "reason", err)
		return
	}
}
