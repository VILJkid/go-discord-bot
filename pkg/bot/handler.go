package bot

import (
	"context"
	"log/slog"
	"strings"

	"github.com/VILJkid/go-discord-bot/pkg/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

type Command interface {
	SetCommandConfig(context.Context, *discordgo.Session, *discordgo.MessageCreate, []string) error
	ValidateCommand(context.Context) error
	ExecuteCommand(context.Context) error
}

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate, c string, a []string) {
	var commandStruct Command

	switch strings.ToLower(c) {
	case "ping":
		commandStruct = new(commands.Ping)
	case "pong":
		commandStruct = new(commands.Pong)
	default:
		slog.Error("unknown command provided:", "command", c)
		return
	}

	type altString string
	var r altString = "requestId"

	ctx := context.WithValue(context.Background(), r, uuid.New().String())
	if err := commandStruct.SetCommandConfig(ctx, s, m, a); err != nil {
		slog.ErrorContext(ctx, "unable to set the command config:", "command", c, "reason", err)
		return
	}
	if err := commandStruct.ValidateCommand(ctx); err != nil {
		slog.ErrorContext(ctx, "unable to validate the command:", "command", c, "reason", err)
		return
	}
	if err := commandStruct.ExecuteCommand(ctx); err != nil {
		slog.ErrorContext(ctx, "unable to execute the command:", "command", c, "reason", err)
		return
	}
}
