package commands

import (
	"context"
	"log"

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
		Label: label,
		Style: discordgo.SuccessButton,
		// Emoji:    discordgo.ComponentEmoji{},
		CustomID: utils.InteractionCustomIDCreateChannelButton,
	}

	a := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{button},
			},
		},
	}

	sentMessage, err := s.ChannelMessageSendComplex(m.ChannelID, a)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	log.Println("Message sent with button. Message ID:", sentMessage.ID)
	return
}
