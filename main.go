package main

import (
	"embed"
	"io/fs"
	"log"
	"minecraftevent/database"
	"minecraftevent/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	// Initialiser la base de données
	if err := database.InitDB(); err != nil {
		log.Fatal("Erreur d'initialisation de la base de données:", err)
	}

	// Seed data (événements de démonstration)
	if err := database.SeedData(); err != nil {
		log.Println("Erreur lors du seed:", err)
	}

	// Créer l'instance Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Servir les fichiers statiques avec embed.FS
	assetsSubFS, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		log.Fatal("Erreur lors de la création du sous-système de fichiers:", err)
	}
	e.StaticFS("/assets", assetsSubFS)

	// Routes
	e.GET("/", handlers.IndexHandler)
	e.GET("/create", handlers.CreateEventPageHandler)
	e.POST("/create", handlers.CreateEventHandler)
	e.GET("/edit/:id", handlers.EditEventPageHandler)
	e.POST("/edit/:id", handlers.EditEventHandler)
	e.DELETE("/delete/:id", handlers.DeleteEventHandler)

	// Démarrer le serveur
	log.Println("🚀 Serveur démarré sur http://localhost:3000")
	e.Logger.Fatal(e.Start(":3000"))
}
