// Package bot creates functions used by discordgo.
package bot

import "github.com/bwmarrin/discordgo"

// MessageCreateHandler handles received messages from channels.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
