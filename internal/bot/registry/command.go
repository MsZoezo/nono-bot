package registry

import (
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

// Command interface that defines the required methods for our commands
type Command interface {
	Run(db *db.Database, s *discordgo.Session, i *discordgo.InteractionCreate) error

	Definition() *discordgo.ApplicationCommand
}
