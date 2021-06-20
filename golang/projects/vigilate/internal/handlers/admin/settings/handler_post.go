package settings

import (
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
)

type HandlerPost struct {
	renderer *services.Renderer
	prefs    *helpers.Preferences
	repo     Repository
	session  *session.Session
}

func NewHandlerPost(renderer *services.Renderer, prefs *helpers.Preferences, repo Repository, session *session.Session) *HandlerPost {
	return &HandlerPost{renderer: renderer, prefs: prefs, repo: repo, session: session}
}

// ServeHTTP saves site settings
func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefMap := make(map[string]string)

	prefMap["site_url"] = r.Form.Get("site_url")
	prefMap["notify_name"] = r.Form.Get("notify_name")
	prefMap["notify_email"] = r.Form.Get("notify_email")
	prefMap["smtp_server"] = r.Form.Get("smtp_server")
	prefMap["smtp_port"] = r.Form.Get("smtp_port")
	prefMap["smtp_user"] = r.Form.Get("smtp_user")
	prefMap["smtp_password"] = r.Form.Get("smtp_password")
	prefMap["sms_enabled"] = r.Form.Get("sms_enabled")
	prefMap["sms_provider"] = r.Form.Get("sms_provider")
	prefMap["twilio_phone_number"] = r.Form.Get("twilio_phone_number")
	prefMap["twilio_sid"] = r.Form.Get("twilio_sid")
	prefMap["twilio_auth_token"] = r.Form.Get("twilio_auth_token")
	prefMap["smtp_from_email"] = r.Form.Get("smtp_from_email")
	prefMap["smtp_from_name"] = r.Form.Get("smtp_from_name")
	prefMap["notify_via_sms"] = r.Form.Get("notify_via_sms")
	prefMap["notify_via_email"] = r.Form.Get("notify_via_email")
	prefMap["sms_notify_number"] = r.Form.Get("sms_notify_number")

	if r.Form.Get("sms_enabled") == "0" {
		prefMap["notify_via_sms"] = "0"
	}

	err := h.repo.InsertOrUpdateSitePreferences(prefMap)
	if err != nil {
		log.Println(err)
		h.renderer.ClientError(w, r, http.StatusBadRequest)
		return
	}

	// update app config
	for k, v := range prefMap {
		h.prefs.SetPref(k, v)
	}

	h.session.Manager.Put(r.Context(), "flash", "Changes saved")

	if r.Form.Get("action") == "1" {
		http.Redirect(w, r, "/admin/overview", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/settings", http.StatusSeeOther)
	}
}
