package render

import (
	"net/http"
	"testing"

	"github.com/enesanbar/workspace/projects/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	request, err := getSession()
	if err != nil {
		t.Error(request)
	}

	session.Put(request.Context(), "flash", "123")
	result := AddDefaultData(request, &td)
	if result.Flash != "123" {
		t.Errorf("expected 123, but got %s", result.Flash)
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	cache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = cache

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}
	writer := myWriter{}
	err = Template(writer, request, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = Template(writer, request, "non-existing.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}
	ctx := request.Context()
	ctx, _ = session.Load(ctx, request.Header.Get("X-Session"))
	request = request.WithContext(ctx)
	return request, err
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
