package schedule

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	GetHostServiceByID(id int) (models.HostService, error)
}
