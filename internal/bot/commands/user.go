package commands

import (
	"errors"
	"fmt"
	"strconv"
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
	member, _ := s.GuildMember(i.GuildID, user.ID)

	if s.State.User.ID == user.ID {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey, what are you inspecting **me** for?",
			},
		})

		return nil
	}

	if user == nil {
		log.Error("Error getting user option.")
		return errors.New("Required option user missing")
	}

	words, events, err := u.Db.GetUserInfo(user.ID, i.GuildID)

	if err != nil {
		log.Error("Error getting top words.")
		return err
	}

	var wsb strings.Builder

	if len(words) == 0 {
		wsb.WriteString("Wow squeeky clean..")
	}

	for i, o := range words {
		wsb.WriteString(fmt.Sprintf("%d. **%s** → %d\n", i+1, o.Word, o.Total))
	}

	var esb strings.Builder

	if len(events) == 0 {
		esb.WriteString("Suspiciously clean..")
	}

	for _, e := range events {
		t := e.Timestamp

		esb.WriteString(fmt.Sprintf("**%s/%02d/%02d %02d:%02d** → %s\n", strconv.Itoa(t.Year())[2:], t.Month(), t.Day(), t.Hour(), t.Minute(), e.Word))
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Inspecting %s", member.Nick),
					Image: &discordgo.MessageEmbedImage{URL: user.AvatarURL("128")},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Top words",
							Value: wsb.String(),
						},
						{
							Name:  "Latest infractions",
							Value: esb.String(),
						},
					},
					Color: member.User.AccentColor,
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
