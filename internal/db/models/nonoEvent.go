// Package models contains all database models.
package models

import (
	"time"

	"gorm.io/gorm"
)

// NonoEvent is a model containing an event of a nono word said in a server.
type NonoEvent struct {
	gorm.Model

	GuildID string
	UserID  string
	Word    string

	Timestamp time.Time
}
