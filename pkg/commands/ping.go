package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Ping struct {
	Context       context.Context
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (p *Ping) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	p.Context = ctx
	p.Session = s
	p.MessageCreate = m
	p.Args = a
	return
}

func (p *Ping) ValidateCommand() (err error) {
	// TODO validation
	return
}

func (p *Ping) ExecuteCommand() (err error) {
	s := p.Session
	m := p.MessageCreate

	successMsg := "Pong!"
	s.ChannelMessageSend(m.ChannelID, successMsg)
	slog.InfoContext(p.Context, "Response sent:", "message", successMsg)
	return
}
