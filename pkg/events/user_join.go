package events

import (
	"github.com/bwmarrin/discordgo"
)

func HandleUserJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// You can access information about the user who joined and the server (guild) they joined.
	user := m.Member.User
	guildID := m.GuildID

	// Perform actions in response to a user joining
	// For example, send a welcome message, assign roles, or log the event.

	// Example: Sending a welcome message
	welcomeMessage := "Welcome to the server, " + user.Username + "!"
	s.ChannelMessageSend(guildID, welcomeMessage)
}
