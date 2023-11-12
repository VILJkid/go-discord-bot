package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Pong struct {
	Context       context.Context
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (p *Pong) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	p.Context = ctx
	p.Session = s
	p.MessageCreate = m
	p.Args = a
	return
}

func (p *Pong) ValidateCommand() (err error) {
	// TODO validation
	return
}

func (p *Pong) ExecuteCommand() (err error) {
	s := p.Session
	m := p.MessageCreate

	successMsg := "Ping!"
	s.ChannelMessageSend(m.ChannelID, successMsg)
	slog.InfoContext(p.Context, "Response sent:", "message", successMsg)
	return
}
