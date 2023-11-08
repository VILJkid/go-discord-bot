package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Ping struct {
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (p *Ping) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	p.Session = s
	p.MessageCreate = m
	p.Args = a
	return
}

func (p *Ping) ValidateCommand(ctx context.Context) (err error) {
	// TODO validation
	return
}

func (p *Ping) ExecuteCommand(ctx context.Context) (err error) {
	s := p.Session
	m := p.MessageCreate

	successMsg := "Pong!"
	s.ChannelMessageSend(m.ChannelID, successMsg)
	slog.InfoContext(ctx, "Response sent:", "message", successMsg)
	return
}
