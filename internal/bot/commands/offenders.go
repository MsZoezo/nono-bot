package commands

import (
	"fmt"
	"strings"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

// Offenders command
type Offenders struct{}

// Run the offenders command
func (Offenders) Run(db *db.Database, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	offenders, err := db.GetTopOffenders(i.GuildID)

	if err != nil {
		log.Error("Error getting top offenders.")
		return err
	}

	var sb strings.Builder

	sb.WriteString("Top offenders:\n")

	for i, o := range offenders {
		user, _ := s.User(o.UserID)
		sb.WriteString(fmt.Sprintf("%d. %s — %d\n", i+1, user.DisplayName(), o.Total))
	}

	guild, _ := s.Guild(i.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("Top offenders - %s", guild.Name),
					Description: sb.String(),
					Color:       i.Member.User.AccentColor,
				},
			},
		},
	})

	return nil
}

// Definition of ping command
func (Offenders) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "offenders",
		Description: "Get top offender!",
	}
}
