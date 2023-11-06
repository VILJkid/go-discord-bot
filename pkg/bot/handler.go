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

	ctx := context.WithValue(context.Background(), utils.ConstContextRequestID, uuid.New().String())
	slog.InfoContext(ctx, "Request started:", string(utils.ConstContextRequestID), ctx.Value(utils.ConstContextRequestID))
	defer slog.InfoContext(ctx, "Request finished:", string(utils.ConstContextRequestID), ctx.Value(utils.ConstContextRequestID))

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
