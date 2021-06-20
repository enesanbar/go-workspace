package service

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	GetServicesByStatus(status string) ([]models.HostService, error)
	UpdateHostServiceStatus(hostID, serviceID, active int) error
	GetHostServiceByHostIDServiceID(hostID int, serviceID int) (models.HostService, error)
	GetHostByID(id int) (models.Host, error)
	GetHostServiceByID(id int) (models.HostService, error)
	InsertEvent(e models.Event) error
	UpdateHostService(hs models.HostService) error
}
