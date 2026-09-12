// Package bot creates functions used by discordgo.
package bot

import (
	"os"
	"os/signal"
	"syscall"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/bot/commands"
	"github.com/MsZoezo/nono-bot/internal/bot/registry"
	"github.com/bwmarrin/discordgo"
)

// Bot struct containing data and functions for connecting to discord
type Bot struct {
	dg *discordgo.Session

	registry *registry.Registry
}

// New creates an instance of the bot with provided token.
func New(token string) (*Bot, error) {
	dg, err := discordgo.New(token)

	if err != nil {
		log.Error("Couldn't intialize discord session..")
		return nil, err
	}

	log.Debug("Succesfully created discord session.")

	dg.Identify.Intents = discordgo.IntentGuildMessages

	registry := registry.New(
		commands.Ping{},
	)

	dg.AddHandler(CreateMessageCreateHandler())
	dg.AddHandler(ConnectHandler)
	dg.AddHandler(DisconnectHandler)
	dg.AddHandler(registry.OnCommand)

	return &Bot{dg, registry}, nil
}

// RegisterCommands registers all commands globally or for the specified guildID
func (bot *Bot) RegisterCommands(guildID string) error {
	bot.dg.Open()
	defer bot.dg.Close()

	return bot.registry.RegisterCommands(bot.dg, guildID)
}

// Run the bot
func (bot *Bot) Run() error {
	err := bot.dg.Open()

	if err != nil {
		log.Error("Couldn't open connection to discord..")
		return err
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}

// Close the connection to discord.
func (bot *Bot) Close() {
	log.Info("Closing connection to discord.")

	bot.dg.Close()
}
