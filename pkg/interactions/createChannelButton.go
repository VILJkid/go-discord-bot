package interactions

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type CreateChannelButton struct {
	Context           context.Context
	Session           *discordgo.Session
	InteractionCreate *discordgo.InteractionCreate
}

func (ccb *CreateChannelButton) SetInteractionConfig(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) (err error) {
	ccb.Context = ctx
	ccb.Session = s
	ccb.InteractionCreate = i
	return
}

func (ccb *CreateChannelButton) InteractionResponse() (err error) {
	s := ccb.Session
	i := ccb.InteractionCreate

	invokingChannel, err := s.Channel(i.ChannelID)
	if err != nil {
		errMsg := "Failed to fetch the channel info!"
		s.ChannelMessageSend(i.ChannelID, errMsg)
		slog.ErrorContext(ccb.Context, errMsg)
		return
	}

	newChannelName := "private-channel"
	newChannel, err := s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
		Name:     newChannelName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: invokingChannel.ParentID,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    i.Member.User.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages,
			},
			{
				ID:   i.GuildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionViewChannel,
			},
		},
	})

	if err != nil {
		errMsg := "Failed to create the private channel!"
		s.ChannelMessageSend(i.ChannelID, errMsg)
		slog.ErrorContext(ccb.Context, errMsg)
		return
	}

	// Send a confirmation message to the user who clicked the button
	if err = ccb.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Private channel %q created for %s. Channel ID: %s", newChannelName, i.Member.User.Mention(), newChannel.ID),
		},
	}); err != nil {
		errMsg := "Failed to respond to the interaction!"
		s.ChannelMessageSend(i.ChannelID, errMsg)
		slog.ErrorContext(ccb.Context, errMsg)
		return
	}

	successMsg := "Interaction handled for create-channel-button!"
	ccb.Session.ChannelMessageSend(i.ChannelID, successMsg)
	slog.InfoContext(ccb.Context, "Interaction handled:", "message", successMsg)
	return
}
