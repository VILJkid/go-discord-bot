package bot

import (
	"context"
	"log/slog"

	"github.com/VILJkid/go-discord-bot/pkg/interactions"
	"github.com/VILJkid/go-discord-bot/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

type Interaction interface {
	SetInteractionConfig(context.Context, *discordgo.Session, *discordgo.InteractionCreate) error
	InteractionResponse() error
}

func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx := context.Background()
	ctx = utils.SetContextRequestID(ctx, uuid.NewString())

	slog.InfoContext(ctx, "Request started:", utils.ConstContextRequestID, utils.GetContextRequestID(ctx))
	defer slog.InfoContext(ctx, "Request finished:", utils.ConstContextRequestID, utils.GetContextRequestID(ctx))

	var interaction Interaction
	switch i.Type {
	case discordgo.InteractionMessageComponent: // Check if the interaction is a button click
		switch i.MessageComponentData().CustomID { // Check if the custom ID of the clicked button matches our expected value
		case utils.InteractionCustomIDCreateChannelButton: // Extract the user who clicked the button
			interaction = new(interactions.CreateChannelButton)
		default:
			return
		}
	default:
		return
	}

	if err := interaction.SetInteractionConfig(ctx, s, i); err != nil {
		// slog.ErrorContext(ctx, "Unable to set the command config:", "command", c, "reason", err)
		return
	}
	if err := interaction.InteractionResponse(); err != nil {
		// slog.ErrorContext(ctx, "Unable to execute the command:", "command", c, "reason", err)
		return
	}
}
