package repository

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

// DatabaseRepo is the database repository
type DatabaseRepo interface {
	// preferences
	AllPreferences() ([]models.Preference, error)
	SetSystemPref(name, value string) error
	UpdateSystemPref(name, value string) error
	InsertOrUpdateSitePreferences(pm map[string]string) error

	// users and authentication
	GetUserById(id int) (models.User, error)
	InsertUser(u models.User) (int, error)
	UpdateUser(u models.User) error
	DeleteUser(id int) error
	UpdatePassword(id int, newPassword string) error
	Authenticate(email, testPassword string) (int, string, error)
	AllUsers() ([]*models.User, error)
	InsertRememberMeToken(id int, token string) error
	DeleteToken(token string) error
	CheckForToken(id int, token string) bool

	InsertHost(host models.Host) (int, error)
	GetHosts() ([]*models.Host, error)
	GetHostByID(id int) (models.Host, error)
	UpdateHost(host models.Host) error

	UpdateHostService(hs models.HostService) error
	UpdateHostServiceStatus(hostID, serviceID, active int) error
	GetAllServiceStatusCounts() (int, int, int, int, error)
	GetServicesByStatus(status string) ([]models.HostService, error)
	GetHostServiceByID(id int) (models.HostService, error)
	GetServicesToMonitor() ([]models.HostService, error)
	GetHostServiceByHostIDServiceID(hostID int, serviceID int) (models.HostService, error)

	InsertEvent(e models.Event) error
	GetAllEvents() ([]models.Event, error)
}
