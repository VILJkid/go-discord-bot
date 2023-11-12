package interactions

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type CreateChannelButton struct {
	Context           context.Context
	Session           *discordgo.Session
	InteractionCreate *discordgo.InteractionCreate
}

// SetCommandConfig and other methods...

func (ccb *CreateChannelButton) SetInteractionConfig(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) (err error) {
	ccb.Context = ctx
	ccb.Session = s
	ccb.InteractionCreate = i
	return
}

func (ccb *CreateChannelButton) InteractionResponse() (err error) {
	// Your logic to handle the interaction for the create-channel-button command
	// ...

	user := ccb.InteractionCreate.Member.User

	// Get the guild ID where the interaction occurred
	guildID := ccb.InteractionCreate.GuildID

	invokingChannel, err := ccb.Session.Channel(ccb.InteractionCreate.ChannelID)
	if err != nil {
		return
	}

	// Create a new channel with the user as the only member
	newChannel, err := ccb.Session.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:     "private-channel",
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: invokingChannel.ParentID,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			// Allow the user to view and send messages
			{
				ID:    user.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages,
			},
			// Deny everyone else the permission to view the channel
			{
				ID:   guildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionViewChannel,
			},
		},
	})

	if err != nil {
		log.Println("Error creating private channel:", err)
		return
	}

	// Send a confirmation message to the user who clicked the button
	err = ccb.Session.InteractionRespond(ccb.InteractionCreate.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Private channel created for %s. Channel ID: %s", user.Mention(), newChannel.ID),
		},
	})

	if err != nil {
		log.Println("Error responding to interaction:", err)
		return
	}

	log.Printf("Private channel created for %s. Channel ID: %s\n", user.Username, newChannel.ID)

	successMsg := "Interaction handled for create-channel-button!"
	ccb.Session.ChannelMessageSend(ccb.InteractionCreate.ChannelID, successMsg)
	slog.InfoContext(ccb.Context, "Interaction handled:", "message", successMsg)
	return
}
