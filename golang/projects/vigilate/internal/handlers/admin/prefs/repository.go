package prefs

type Repository interface {
	UpdateSystemPref(name, value string) error
}
