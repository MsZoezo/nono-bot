// Package registry handles our collection of commands.
package registry

import (
	"github.com/bwmarrin/discordgo"
)

// Registry for our commands
type Registry struct {
	commands map[string]Command
}

// RegisterCommands registers all commands in the register
func (registry *Registry) RegisterCommands(s *discordgo.Session, guildID string) error {
	defs := make([]*discordgo.ApplicationCommand, 0, len(registry.commands))

	for _, command := range registry.commands {
		defs = append(defs, command.Definition())
	}

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildID, defs)

	return err
}

// New registry with these commands
func New(commands ...Command) (registry *Registry) {
	registry = &Registry{commands: make(map[string]Command, len(commands))}

	for _, command := range commands {
		registry.commands[command.Definition().Name] = command
	}

	return
}

// OnCommand handles incoming interactions and runs the corresponding command.
func (registry *Registry) OnCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := registry.commands[i.ApplicationCommandData().Name]; ok {
		cmd.Run(s, i)
	}
}
