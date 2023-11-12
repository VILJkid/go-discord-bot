package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type CreateChannel struct {
	Context       context.Context
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (c *CreateChannel) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	c.Context = ctx
	c.Session = s
	c.MessageCreate = m
	c.Args = a
	return
}

func (c *CreateChannel) ValidateCommand() (err error) {
	// TODO validation
	return
}

func (c *CreateChannel) ExecuteCommand() (err error) {
	s := c.Session
	m := c.MessageCreate

	channelName := "private"
	if len(c.Args) >= 1 {
		channelName = c.Args[0]
	}

	invokingChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		errMsg := "Failed to fetch the channel info!"
		s.ChannelMessageSend(m.ChannelID, errMsg)
		slog.ErrorContext(c.Context, errMsg)
		return
	}

	channel, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     channelName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: invokingChannel.ParentID,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    m.Author.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionViewChannel | discordgo.PermissionManageMessages,
			},
			{
				ID:   m.GuildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionViewChannel,
			},
		},
	})
	if err != nil {
		errMsg := "Failed to create a private channel!"
		s.ChannelMessageSend(m.ChannelID, errMsg)
		slog.ErrorContext(c.Context, errMsg)
		return
	}

	successMsg := channelName + " channel created!"
	s.ChannelMessageSend(m.ChannelID, successMsg)
	slog.InfoContext(c.Context, "Response sent:", "message", successMsg)
	successMsg = "Hi " + m.Author.Mention() + "! Only you can see this channel"
	s.ChannelMessageSend(channel.ID, successMsg)
	slog.InfoContext(c.Context, "Response sent:", "message", successMsg)
	return
}
