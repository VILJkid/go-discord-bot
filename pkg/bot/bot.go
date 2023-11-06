package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	commandPrefix := "!"
	m.Content = strings.TrimSpace(m.Content)
	if !strings.HasPrefix(m.Content, commandPrefix) {
		return
	}

	parts := strings.Fields(m.Content)
	if len(parts) < 1 {
		return
	}

	command := strings.TrimPrefix(parts[0], commandPrefix)
	if len(command) == 0 {
		return
	}

	HandleCommand(s, m, command, parts[1:])
}
