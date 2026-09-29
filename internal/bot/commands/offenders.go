package commands

import (
	"fmt"
	"strings"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

// Offenders command
type Offenders struct {
	Db *db.Database
}

// Run the offenders command
func (o Offenders) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	offenders, err := o.Db.GetTopOffenders(i.GuildID)

	if err != nil {
		log.Error("Error getting top words.")
		return err
	}

	var sb strings.Builder

	for idx, o := range offenders {
		member, _ := s.GuildMember(i.GuildID, o.UserID)
		sb.WriteString(fmt.Sprintf("%d. **%s** → %d\n", idx+1, member.Nick, o.Total))
	}

	guild, _ := s.Guild(i.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Inspecting %s", guild.Name),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Top offenders",
							Value: sb.String(),
						},
					},
					Color: i.Member.User.AccentColor,
				},
			},
		},
	})

	return nil
}

// Definition of offenders command
func (Offenders) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "offenders",
		Description: "Get top offender!",
	}
}
