package commands

import (
	"context"
	"log/slog"

	"github.com/VILJkid/go-discord-bot/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CreateChannelButton struct {
	Context       context.Context
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (c *CreateChannelButton) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	c.Context = ctx
	c.Session = s
	c.MessageCreate = m
	c.Args = a
	return
}

func (c *CreateChannelButton) ValidateCommand() (err error) {
	// TODO validation
	return
}

func (c *CreateChannelButton) ExecuteCommand() (err error) {
	s := c.Session
	m := c.MessageCreate

	label := "Create Channel"
	button := discordgo.Button{
		Label:    label,
		Style:    discordgo.SuccessButton,
		CustomID: utils.InteractionCustomIDCreateChannelButton,
	}

	messageComponentData := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{button},
			},
		},
	}

	s.ChannelMessageSendComplex(m.ChannelID, messageComponentData)
	slog.InfoContext(c.Context, "Message sent with button")
	return
}
