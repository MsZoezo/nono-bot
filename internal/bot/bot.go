// Package bot creates functions used by discordgo.
package bot

import (
	"charm.land/log/v2"
	"github.com/bwmarrin/discordgo"
)

// Bot struct containing data and functions for connecting to discord
type Bot struct {
	dg *discordgo.Session
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

	dg.AddHandler(MessageCreateHandler)
	dg.AddHandler(ConnectHandler)
	dg.AddHandler(DisconnectHandler)

	err = dg.Open()

	if err != nil {
		log.Error("Couldn't open connection to discord..")
		return nil, err
	}

	return &Bot{dg}, nil
}

// Close the connection to discord.
func (bot *Bot) Close() {
	log.Info("Closing connection to discord.")

	bot.dg.Close()
}
