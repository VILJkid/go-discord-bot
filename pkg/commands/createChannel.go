package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type CreateChannel struct {
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (c *CreateChannel) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	c.Session = s
	c.MessageCreate = m
	c.Args = a
	return
}

func (c *CreateChannel) ValidateCommand(ctx context.Context) (err error) {
	// TODO validation
	return
}

func (c *CreateChannel) ExecuteCommand(ctx context.Context) (err error) {
	s := c.Session
	m := c.MessageCreate
	user := m.Author
	guildID := m.GuildID

	channelName := "private-channel"
	if len(c.Args) >= 1 {
		channelName = c.Args[0]
	}

	channel, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name: channelName,
		Type: discordgo.ChannelTypeGuildText,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    user.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionViewChannel | discordgo.PermissionManageMessages,
			},
			{
				ID:   guildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionViewChannel,
			},
		},
	})
	if err != nil {
		errMsg := "Failed to create a private channel!"
		s.ChannelMessageSend(m.ChannelID, errMsg)
		slog.ErrorContext(ctx, errMsg)
		return
	}

	successMsg := "Private channel created!"
	s.ChannelMessageSend(m.ChannelID, successMsg)
	slog.InfoContext(ctx, "Response sent:", "message", successMsg)
	successMsg = "Hi " + m.Author.Mention() + "! Only you can see this channel"
	s.ChannelMessageSend(channel.ID, successMsg)
	slog.InfoContext(ctx, "Response sent:", "message", successMsg)
	return
}
