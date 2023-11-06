package commands

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Pong struct {
	Session       *discordgo.Session
	MessageCreate *discordgo.MessageCreate
	Args          []string
}

func (p *Pong) SetCommandConfig(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, a []string) (err error) {
	p.Session = s
	p.MessageCreate = m
	p.Args = a
	return
}

func (p *Pong) ValidateCommand(ctx context.Context) (err error) {
	// TODO validation
	return
}

func (p *Pong) ExecuteCommand(ctx context.Context) (err error) {
	s := p.Session
	m := p.MessageCreate

	slog.InfoContext(ctx, "Command recieved:", "command", "pong")
	s.ChannelMessageSend(m.ChannelID, "Ping!")
	slog.InfoContext(ctx, "Response sent:", "message", "Ping!")
	return
}
