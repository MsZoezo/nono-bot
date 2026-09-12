package registry

import "github.com/bwmarrin/discordgo"

// Command interface that defines the required methods for our commands
type Command interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error

	Definition() *discordgo.ApplicationCommand
}
