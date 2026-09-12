// Package commands contains every command for our discord bot
package commands

import "github.com/bwmarrin/discordgo"

// Ping command
type Ping struct{}

// Run the ping command
func (ping Ping) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})

	return nil
}

// Definition of ping command
func (ping Ping) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot!",
	}
}
