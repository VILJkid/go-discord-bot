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
	ValidateCommand(context.Context) error
	ExecuteCommand(context.Context) error
}

func ErrorMessageEmbed(description string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: description,
		Color:       16711680,
	}
}

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate, c string, a []string) {
	var commandStruct Command

	ctx := context.WithValue(context.Background(), utils.ConstContextRequestID, uuid.New().String())
	slog.InfoContext(ctx, "Request started:", string(utils.ConstContextRequestID), ctx.Value(utils.ConstContextRequestID))
	defer slog.InfoContext(ctx, "Request finished:", string(utils.ConstContextRequestID), ctx.Value(utils.ConstContextRequestID))

	c = strings.ToLower(c)
	switch c {
	case utils.CommandPing:
		commandStruct = new(commands.Ping)
	case utils.CommandPong:
		commandStruct = new(commands.Pong)
	case utils.CommandCreateChannel:
		commandStruct = new(commands.CreateChannel)
	default:
		embed := ErrorMessageEmbed("Command not recognized: " + c)
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		slog.Error("Command not recognized:", "command", c)
		return
	}

	slog.InfoContext(ctx, "Command recieved:", "command", c)

	if err := commandStruct.SetCommandConfig(ctx, s, m, a); err != nil {
		slog.ErrorContext(ctx, "Unable to set the command config:", "command", c, "reason", err)
		return
	}
	if err := commandStruct.ValidateCommand(ctx); err != nil {
		slog.ErrorContext(ctx, "Unable to validate the command:", "command", c, "reason", err)
		return
	}
	if err := commandStruct.ExecuteCommand(ctx); err != nil {
		slog.ErrorContext(ctx, "Unable to execute the command:", "command", c, "reason", err)
		return
	}
}
