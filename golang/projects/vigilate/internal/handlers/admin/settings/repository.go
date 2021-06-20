package settings

type Repository interface {
	InsertOrUpdateSitePreferences(pm map[string]string) error
}
