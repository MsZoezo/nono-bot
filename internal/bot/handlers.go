// Package bot creates functions used by discordgo.
package bot

import (
	"charm.land/log/v2"
	"github.com/bwmarrin/discordgo"
)

// MessageCreateHandler handles received messages from channels.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Debug("Received new message.", "User", m.Author.Username)

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

// ConnectHandler handles state after connecting to discord.
func ConnectHandler(s *discordgo.Session, _ *discordgo.Connect) {
	log.Info("Succesfully connected.", "Shard", s.ShardID)
}

// DisconnectHandler handles state after being disconnected from discord.
func DisconnectHandler(_ *discordgo.Session, _ *discordgo.Disconnect) {
	log.Fatal("Disconnected from discord..")
}
