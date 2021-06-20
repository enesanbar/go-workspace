package services

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
	"github.com/justinas/nosurf"
)

type Renderer struct {
	Session *session.Session
	Prefs   *helpers.Preferences
}

func NewRenderer(session *session.Session, prefs *helpers.Preferences) *Renderer {
	return &Renderer{Session: session, Prefs: prefs}
}

// DefaultData adds default data which is accessible to all templates
func (ren *Renderer) DefaultData(td models.TemplateData, r *http.Request, w http.ResponseWriter) models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	td.IsAuthenticated = ren.Session.IsAuthenticated(r)
	td.PreferenceMap = ren.Prefs.Prefs
	// if logged in, store user id in template data
	if td.IsAuthenticated {
		u := ren.Session.Manager.Get(r.Context(), "user").(models.User)
		td.User = u
	}

	td.Flash = ren.Session.Manager.PopString(r.Context(), "flash")
	td.Warning = ren.Session.Manager.PopString(r.Context(), "warning")
	td.Error = ren.Session.Manager.PopString(r.Context(), "error")

	return td
}

// RenderPage renders a page using jet templates
func (ren *Renderer) RenderPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	// add default template data
	var td models.TemplateData
	if data != nil {
		td = data.(models.TemplateData)
	}

	// add default data
	td = ren.DefaultData(td, r, w)

	// add template functions
	addTemplateFunctions()

	// load the template and render it
	t, err := views.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ren *Renderer) PrintTemplateError(w http.ResponseWriter, err error) {
	_, _ = fmt.Fprint(w, fmt.Sprintf(`<small><span class='text-danger'>Error executing template: %s</span></small>`, err))

}

func (ren *Renderer) ClientError(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		show404(w, r)
	case http.StatusInternalServerError:
		show500(w, r)
	default:
		http.Error(w, http.StatusText(status), status)
	}
}

// ServerError will display error page for internal server error
func (ren *Renderer) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = log.Output(2, trace)

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	http.ServeFile(w, r, "./ui/static/500.html")
}

func show404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	http.ServeFile(w, r, "./ui/static/404.html")
}

func show500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	http.ServeFile(w, r, "./ui/static/500.html")
}
