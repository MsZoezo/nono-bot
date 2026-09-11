package bot

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/text"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// CreateMessageCreateHandler creates the function that handles received messages from channels.
func CreateMessageCreateHandler() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	words := viper.GetStringSlice("filter.default")

	var filter text.Filter

	if len(words) != 0 {
		filter = text.ArrToFilter(words)
	} else {
		log.Info("No default filter list found in config.")
	}

	responses := viper.GetStringSlice("filter.responses")

	if len(responses) == 0 {
		log.Fatal("No responses found in config..")
	}

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if len(filter) == 0 {
			return
		}

		if m.Author.ID == s.State.User.ID {
			return
		}

		start := time.Now()

		found := text.ContainsBadWords(filter, m.Content)

		if len(found) == 0 {
			return
		}

		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)

		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(responses))))

		response := responses[random.Int64()]

		msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(response, m.Author.Mention()))

		end := time.Now()

		elapsed := end.Sub(start)

		log.Debug("Filtered out bad words!", "user", m.Author.DisplayName(), "Elapsed (ms)", elapsed.Milliseconds(), "words", found)

		time.Sleep(5 * time.Second)

		s.ChannelMessageDelete(msg.ChannelID, msg.ID)
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
