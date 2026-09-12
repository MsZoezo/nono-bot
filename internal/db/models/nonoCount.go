// Package models contains all database models.
package models

import "gorm.io/gorm"

// NonoCount is a model containing how often a user in a server said a nono word.
type NonoCount struct {
	gorm.Model

	GuildID string `gorm:"uniqueIndex:idx_guild_user_word"`
	UserID  string `gorm:"uniqueIndex:idx_guild_user_word"`
	Word    string `gorm:"uniqueIndex:idx_guild_user_word"`

	Count uint64
}
