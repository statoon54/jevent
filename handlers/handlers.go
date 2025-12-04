package handlers

import (
	"minecraftevent/database"
	"minecraftevent/models"
	"minecraftevent/templates"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Handler pour la page d'accueil affichant la liste des événements
func IndexHandler(c echo.Context) error {
	var events []models.Event
	result := database.DB.Order("start_date DESC").Find(&events)
	if result.Error != nil {
		return c.String(
			http.StatusInternalServerError,
			"Erreur lors de la récupération des événements",
		)
	}

	return templates.Index(events).Render(c.Request().Context(), c.Response().Writer)
}

// Handlers pour la création, l'édition et la suppression des événements
func CreateEventPageHandler(c echo.Context) error {
	return templates.CreateEvent().Render(c.Request().Context(), c.Response().Writer)
}

// Handler pour créer un nouvel événement
func CreateEventHandler(c echo.Context) error {
	title := c.FormValue("title")
	organizer := c.FormValue("organizer")
	description := c.FormValue("description")
	startDateStr := c.FormValue("start_date")
	endDateStr := c.FormValue("end_date")
	imageURL := c.FormValue("image_url")
	discordURL := c.FormValue("discord_url")

	if title == "" || organizer == "" || startDateStr == "" {
		return c.String(
			http.StatusBadRequest,
			"Les champs titre, organisateur et date de début sont requis",
		)
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Format de date de début invalide")
	}

	var endDate *time.Time
	if endDateStr != "" {
		parsed, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Format de date de fin invalide")
		}
		endDate = &parsed
	}

	event := models.Event{
		Title:       title,
		Organizer:   organizer,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		ImageURL:    imageURL,
		DiscordURL:  discordURL,
	}

	result := database.DB.Create(&event)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Erreur lors de la création de l'événement")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

// Handler pour afficher la page d'édition d'un événement
func EditEventPageHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "ID invalide")
	}

	var event models.Event
	result := database.DB.First(&event, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Événement non trouvé")
	}

	return templates.EditEvent(event).Render(c.Request().Context(), c.Response().Writer)
}

// Handler pour mettre à jour un événement existant
func EditEventHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "ID invalide")
	}

	var event models.Event
	result := database.DB.First(&event, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Événement non trouvé")
	}

	title := c.FormValue("title")
	organizer := c.FormValue("organizer")
	description := c.FormValue("description")
	startDateStr := c.FormValue("start_date")
	endDateStr := c.FormValue("end_date")
	imageURL := c.FormValue("image_url")
	discordURL := c.FormValue("discord_url")

	if title == "" || organizer == "" || startDateStr == "" {
		return c.String(
			http.StatusBadRequest,
			"Les champs titre, organisateur et date de début sont requis",
		)
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Format de date de début invalide")
	}

	var endDate *time.Time
	if endDateStr != "" {
		parsed, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Format de date de fin invalide")
		}
		endDate = &parsed
	}

	event.Title = title
	event.Organizer = organizer
	event.Description = description
	event.StartDate = startDate
	event.EndDate = endDate
	event.ImageURL = imageURL
	event.DiscordURL = discordURL

	result = database.DB.Save(&event)
	if result.Error != nil {
		return c.String(
			http.StatusInternalServerError,
			"Erreur lors de la mise à jour de l'événement",
		)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

// Handler pour supprimer un événement
func DeleteEventHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "ID invalide")
	}

	result := database.DB.Delete(&models.Event{}, id)
	if result.Error != nil {
		return c.String(
			http.StatusInternalServerError,
			"Erreur lors de la suppression de l'événement",
		)
	}

	return c.NoContent(http.StatusOK)
}
