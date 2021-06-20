package hosts

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	GetHosts() ([]*models.Host, error)
	GetHostByID(id int) (models.Host, error)
	UpdateHost(host models.Host) error
	InsertHost(host models.Host) (int, error)
}
