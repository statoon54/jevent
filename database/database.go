package database

import (
	"log"
	"time"

	"minecraftevent/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initialise la connexion à la base de données et effectue les migrations
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Event{})
	if err != nil {
		return err
	}

	log.Println("Database initialized and migrated successfully")
	return nil
}

// SeedData insère des données de démonstration dans la base de données si elle est vide
func SeedData() error {
	var count int64
	DB.Model(&models.Event{}).Count(&count)
	if count > 0 {
		return nil
	}

	events := []models.Event{
		{
			Title:       "Naissance Matthieu",
			Organizer:   "admin",
			Description: "Bah c'est la naissance de la personne la plus importante de l'univers",
			StartDate:   time.Date(1973, 1, 23, 0, 0, 0, 0, time.UTC),
			DiscordURL:  "https://discord.gg/e3w25cnA",
			ImageURL:    "https://images.unsplash.com/photo-1530103862676-de8c9debad1d",
		},
		{
			Title:       "Lancement de Linux",
			Organizer:   "Microsoft",
			Description: "Le système d'exploitation révolutionnaire",
			StartDate:   time.Date(1991, 9, 17, 0, 0, 0, 0, time.UTC),
			ImageURL:    "https://images.unsplash.com/photo-1629654297299-c8506221ca97",
		},
		{
			Title:       "Lancement de Windows 95",
			Organizer:   "admin",
			Description: "Une révolution dans le monde de l'informatique",
			StartDate:   time.Date(1995, 5, 7, 0, 0, 0, 0, time.UTC),
			ImageURL:    "https://images.unsplash.com/photo-1633356122544-f134324a6cee",
		},
	}

	return DB.Create(&events).Error
}
