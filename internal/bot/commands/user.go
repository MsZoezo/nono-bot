package commands

import (
	"errors"
	"fmt"
	"strings"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

// User command
type User struct {
	Db *db.Database
}

// Run the words command
func (u User) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	user := i.ApplicationCommandData().GetOption("user").UserValue(s)

	if user == nil {
		log.Error("Error getting user option.")
		return errors.New("Required option user missing")
	}

	words, err := u.Db.GetUserWords(user.ID)

	if err != nil {
		log.Error("Error getting top words.")
		return err
	}

	var sb strings.Builder

	for i, o := range words {
		sb.WriteString(fmt.Sprintf("%d. %s — %d\n", i+1, o.Word, o.Total))
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("Inspecting %s", user.DisplayName()),
					Image:       &discordgo.MessageEmbedImage{URL: user.AvatarURL("128")},
					Description: sb.String(),
					Color:       user.AccentColor,
				},
			},
		},
	})

	return nil
}

// Definition of user command
func (User) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "inspect",
		Description: "Get user information!",
		Options: []*discordgo.ApplicationCommandOption{
			{Name: "user", Required: true, Description: "The user to inspect.", Type: discordgo.ApplicationCommandOptionUser},
		},
	}
}
