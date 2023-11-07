package commands

import (
	"context"
	"log/slog"

	"github.com/VILJkid/go-discord-bot/pkg/utils"
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

	slog.InfoContext(ctx, "Command recieved:", "command", utils.CommandPing)
	s.ChannelMessageSend(m.ChannelID, "Pong!")
	slog.InfoContext(ctx, "Response sent:", "message", "Pong!")
	return
}
