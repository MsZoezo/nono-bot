package bot

import (
	"fmt"
	"time"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/text"
	"github.com/bwmarrin/discordgo"
)

// MessageCreateHandler handles received messages from channels.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	start := time.Now()

	found := text.ContainsBadWords(text.ArrToFilter([]string{"pickle"}), m.Content)

	if len(found) == 0 {
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, how could you say such a bad word?!", m.Author.Mention()))

	end := time.Now()

	elapsed := end.Sub(start)

	log.Debug("Filtered out bad words!", "user", m.Author.DisplayName(), "Elapsed (ms)", elapsed.Milliseconds(), "words", found)

	time.Sleep(5 * time.Second)

	s.ChannelMessageDelete(msg.ChannelID, msg.ID)
}

// ConnectHandler handles state after connecting to discord.
func ConnectHandler(s *discordgo.Session, _ *discordgo.Connect) {
	log.Info("Succesfully connected.", "Shard", s.ShardID)
}

// DisconnectHandler handles state after being disconnected from discord.
func DisconnectHandler(_ *discordgo.Session, _ *discordgo.Disconnect) {
	log.Fatal("Disconnected from discord..")
}
