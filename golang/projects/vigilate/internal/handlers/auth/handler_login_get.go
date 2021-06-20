package auth

import (
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
)

type HandlerLoginGet struct {
	Session  *session.Session
	Renderer *services.Renderer
}

func NewHandlerLoginGet(session *session.Session, renderer *services.Renderer) *HandlerLoginGet {
	return &HandlerLoginGet{Session: session, Renderer: renderer}
}

// ServeHTTP shows the home (login) screen
func (h *HandlerLoginGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if already logged in, take to dashboard
	if h.Session.Manager.Exists(r.Context(), "userID") {
		http.Redirect(w, r, "/admin/overview", http.StatusSeeOther)
		return
	}

	err := h.Renderer.RenderPage(w, r, "login", nil, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
