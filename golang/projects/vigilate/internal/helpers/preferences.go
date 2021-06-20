package helpers

import (
	"log"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
)

//go:generate mockery --name=PrefRepository
type PrefRepository interface {
	AllPreferences() ([]models.Preference, error)
}

type Preferences struct {
	Prefs      map[string]string
	Repository repository.DatabaseRepo
}

func NewPreferences(repository PrefRepository) *Preferences {
	preferences, err := repository.AllPreferences()
	if err != nil {
		log.Fatal("Cannot read preferences:", err)
	}

	preferenceMap := make(map[string]string)
	for _, pref := range preferences {
		preferenceMap[pref.Name] = string(pref.Preference)
	}

	//preferenceMap["version"] = main.vigilateVersion

	return &Preferences{
		Prefs: preferenceMap,
	}
}

func (p *Preferences) GetPref(key string) string {
	return p.Prefs[key]
}

func (p *Preferences) SetPref(key, value string) {
	p.Prefs[key] = value
}
