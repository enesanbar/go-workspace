package events

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	GetAllEvents() ([]models.Event, error)
}
