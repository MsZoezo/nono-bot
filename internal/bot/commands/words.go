package commands

import (
	"fmt"
	"strings"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

// Words command
type Words struct {
	Db *db.Database
}

// Run the words command
func (w Words) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	words, err := w.Db.GetTopWordsGuild(i.GuildID)

	if err != nil {
		log.Error("Error getting top words.")
		return err
	}

	var sb strings.Builder

	for i, o := range words {
		sb.WriteString(fmt.Sprintf("%d. %s — %d\n", i+1, o.Word, o.Total))
	}

	guild, _ := s.Guild(i.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("Top words - %s", guild.Name),
					Description: sb.String(),
					Color:       i.Member.User.AccentColor,
				},
			},
		},
	})

	return nil
}

// Definition of words command
func (Words) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "words",
		Description: "Get top words!",
	}
}
