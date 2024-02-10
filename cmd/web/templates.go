package main

import "eventurely/internal/models"

type templateData struct {
	Event  *models.Event
	Events []models.Event
}
