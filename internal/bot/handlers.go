package bot

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/text"
	"github.com/bwmarrin/discordgo"
)

var bannedWords = map[string]bool{
	"pickle": true,
}

// MessageCreateHandler handles received messages from channels.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	found := text.ContainsBadWords(bannedWords, m.Content)

	if len(found) == 0 {
		return
	}

	log.Debug("Filtered out bad words!", "user", m.Author.DisplayName(), "words", found)

	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, how could you say such a bad word?!", m.Author.Mention()))
}

// ConnectHandler handles state after connecting to discord.
func ConnectHandler(s *discordgo.Session, _ *discordgo.Connect) {
	log.Info("Succesfully connected.", "Shard", s.ShardID)
}

// DisconnectHandler handles state after being disconnected from discord.
func DisconnectHandler(_ *discordgo.Session, _ *discordgo.Disconnect) {
	log.Fatal("Disconnected from discord..")
}
