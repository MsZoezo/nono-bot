// Package handlers contains the handlers for discord events.
package handlers

import (
	"fmt"
	"math/rand/v2"
	"time"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/MsZoezo/nono-bot/internal/text"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Filter handles filtering messages on discord
type Filter struct {
	db *db.Database

	defaultFilter text.Filter
	responses     []string
}

// NewFilter creates a new filter with default config options
func NewFilter(db *db.Database) (filter *Filter) {
	filter = &Filter{db: db}

	words := viper.GetStringSlice("filter.default")

	if len(words) != 0 {
		filter.defaultFilter = text.ArrToFilter(words)
	} else {
		log.Info("No default filter list found in config.")
	}

	filter.responses = viper.GetStringSlice("filter.responses")

	if len(filter.responses) == 0 {
		log.Error("No responses found in config..")
	}

	return
}

// Handler handles received messages to filter from channels.
func (filter *Filter) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(filter.defaultFilter) == 0 {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	start := time.Now()

	found := text.ContainsBadWords(filter.defaultFilter, m.Content)

	if len(found) == 0 {
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.Message.ID)

	random := rand.IntN(len(filter.responses))

	response := filter.responses[random]

	msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(response, m.Author.Mention()))

	end := time.Now()

	elapsed := end.Sub(start)

	for word, count := range found {
		go filter.db.UpsertNonoWord(m.GuildID, m.Author.ID, word, count)
	}

	log.Debug("Filtered out bad words!", "user", m.Author.DisplayName(), "Elapsed (ms)", elapsed.Milliseconds(), "words", found)

	time.Sleep(5 * time.Second)

	s.ChannelMessageDelete(msg.ChannelID, msg.ID)
}
