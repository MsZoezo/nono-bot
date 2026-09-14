package handlers

import (
	"charm.land/log/v2"
	"github.com/bwmarrin/discordgo"
)

// ConnectHandler handles state after connecting to discord.
func ConnectHandler(s *discordgo.Session, _ *discordgo.Connect) {
	log.Info("Succesfully connected.", "Shard", s.ShardID)
}

// DisconnectHandler handles state after being disconnected from discord.
func DisconnectHandler(_ *discordgo.Session, _ *discordgo.Disconnect) {
	log.Fatal("Disconnected from discord..")
}
