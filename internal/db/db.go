// Package db is a wrapper around the ORM.
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MsZoezo/nono-bot/internal/db/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ctx context.Context

// Database contains all data necessary to work with our orm.
type Database struct {
	db *gorm.DB
}

// New database instance
func New() (*Database, error) {
	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.NonoCount{}, &models.NonoEvent{}); err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// UpsertNonoWord creates or updates a nono word record.
func (database *Database) UpsertNonoWord(guildID string, userID string, word string, count uint64) error {
	row := models.NonoCount{
		GuildID: guildID,
		UserID:  userID,
		Word:    word,
		Count:   count,
	}

	return database.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "guild_id"}, {Name: "user_id"}, {Name: "word"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"count": gorm.Expr("nono_counts.count + ?", row.Count),
		}),
	}).Create(&row).Error
}

// CreateNonoEvent creates a new nono event record.
func (database *Database) CreateNonoEvent(guildID string, userID string, word string) error {
	row := models.NonoEvent{
		GuildID: guildID,
		UserID:  userID,
		Word:    word,

		Timestamp: time.Now(),
	}

	return database.db.Create(&row).Error
}

// OffenderStat is the data returned when we select the top offenders.
type OffenderStat struct {
	UserID string `gorm:"column:user_id"`
	Total  uint64 `gorm:"column:total"`
}

// GetTopOffenders returns the top offenders for a given guild.
func (database *Database) GetTopOffenders(guildID string) ([]OffenderStat, error) {
	var results []OffenderStat

	err := gorm.G[models.NonoCount](database.db).
		Select("user_id, sum(count) AS total").
		Where("guild_id = ?", guildID).
		Group("user_id").Order("total DESC").
		Limit(10).
		Scan(ctx, &results)

	return results, err
}

// WordStat is the data returned when we select the top words.
type WordStat struct {
	Word  string `gorm:"column:word"`
	Total uint64 `gorm:"column:total"`
}

// EventStat is the data returned when we select the latest events.
type EventStat struct {
	Word      string    `gorm:"column:word"`
	UserID    string    `gorm:"column:user_id"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

// GetTopWordsGuild returns the top words for a given guild.
func (database *Database) GetTopWordsGuild(guildID string) ([]WordStat, error) {
	var results []WordStat

	err := gorm.G[models.NonoCount](database.db).
		Select("word, sum(count) AS total").
		Where("guild_id = ?", guildID).
		Group("word").
		Order("total DESC").
		Limit(10).
		Scan(ctx, &results)

	return results, err
}

// GetUserInfo returns nono word counts and events said by a user
func (database *Database) GetUserInfo(userID string, guildID string) ([]WordStat, []EventStat, error) {
	var words []WordStat
	var events []EventStat

	err := gorm.G[models.NonoCount](database.db).
		Select("word, count AS total").
		Where("user_id = ?", userID).
		Order("total DESC").
		Scan(ctx, &words)

	if err != nil {
		return nil, nil, err
	}

	err = gorm.G[models.NonoEvent](database.db).
		Select("word, user_id, timestamp").
		Where("user_id = ?", userID).
		Where("guild_id = ?", guildID).
		Order("timestamp DESC").
		Limit(10).
		Scan(ctx, &events)

	return words, events, err
}
