package models

// TemplateData defines template data
type TemplateData struct {
	CSRFToken       string
	IsAuthenticated bool
	PreferenceMap   map[string]string
	User            User
	Flash           string
	Warning         string
	Error           string
	GwVersion       string
}
