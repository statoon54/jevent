package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `                  json:"created_at"`
	UpdatedAt   time.Time      `                  json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"      json:"-"`
	Title       string         `gorm:"not null"   json:"title"`
	Organizer   string         `gorm:"not null"   json:"organizer"`
	Description string         `                  json:"description"`
	StartDate   time.Time      `gorm:"not null"   json:"start_date"`
	EndDate     *time.Time     `                  json:"end_date"`
	ImageURL    string         `                  json:"image_url"`
	DiscordURL  string         `                  json:"discord_url"`
}
