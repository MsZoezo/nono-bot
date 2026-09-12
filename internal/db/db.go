// Package db is a wrapper around the ORM.
package db

import (
	"context"

	"github.com/MsZoezo/nono-bot/internal/db/models"
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
	db, err := gorm.Open(postgres.Open("host=localhost user=nonobot password=develop dbname=nonodata port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.NonoCount{}); err != nil {
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
		Group("user_id").Order("total DESC").Limit(10).Scan(ctx, &results)

	return results, err
}
